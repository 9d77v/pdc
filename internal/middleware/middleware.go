package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/9d77v/pdc/internal/consts"
	"github.com/9d77v/pdc/internal/db/mq"
	"github.com/9d77v/pdc/internal/module/user-service/models"
	"github.com/9d77v/pdc/internal/utils"
	"github.com/9d77v/pdc/pkg/iot/sdk/pb"
	"github.com/gorilla/websocket"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

//Error ...
type Error struct {
	Message string   `json:"message"`
	Path    []string `json:"path"`
}

//ErrorResponse ..
type ErrorResponse struct {
	Errors []*Error `json:"errors"`
}

//NewErrorResponse ..
func NewErrorResponse(message, path string) *ErrorResponse {
	res := new(ErrorResponse)
	err := &Error{
		Message: message,
		Path:    []string{path},
	}
	res.Errors = []*Error{err}
	return res
}

func resString(errRes *ErrorResponse) string {
	ret, err := json.Marshal(errRes)
	if err != nil {
		log.Println("json marshal faield:", err)
	}
	return string(ret)
}

//Auth ..
func Auth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/api" {
				req, err := parseBody(r)
				if err != nil {
					w.Header().Set("Content-type", "application/json")
					http.Error(w, resString(NewErrorResponse("Invalid request", r.URL.Path)), http.StatusBadRequest)
					return
				}
				//不需要验证的方法
				for _, v := range publicOperationArr {
					if req.OperationName == v {
						next.ServeHTTP(w, r)
						return
					}
				}
				//token验证
				accessToken := r.Header.Get("Authorization")
				if accessToken == "" {
					w.Header().Set("Content-type", "application/json")
					http.Error(w, resString(NewErrorResponse("Invalid token", r.URL.Path)), http.StatusUnauthorized)
					return
				}
				accessToken = strings.TrimPrefix(accessToken, "Bearer ")
				token, err := utils.ParseJWT([]byte(consts.JWTtAccessSecret), accessToken)
				if err != nil {
					w.Header().Set("Content-type", "application/json")
					http.Error(w, resString(NewErrorResponse("Invalid token", r.URL.Path)), http.StatusUnauthorized)
					return
				}
				data, _ := token.Claims.(*utils.MyCustomClaims)
				if data.Issuer != consts.JWTIssuer {
					w.Header().Set("Content-type", "application/json")
					http.Error(w, resString(NewErrorResponse("Invalid token", r.URL.Path)), http.StatusUnauthorized)
					return
				}
				//获取hash前用户id
				uid := consts.GetDecodeUID(data.UID)
				user, err := userService.GetByID(r.Context(), int64(uid))
				if err != nil {
					w.Header().Set("Content-type", "application/json")
					http.Error(w, resString(NewErrorResponse("Invalid token", r.URL.Path)), http.StatusUnauthorized)
					return
				}
				//权限校验
				roleArr := permissonMap[req.OperationName]
				valid := false
				for _, v := range roleArr {
					if v == user.RoleID {
						valid = true
						break
					}
				}
				if !valid {
					w.Header().Set("Content-type", "application/json")
					http.Error(w, resString(NewErrorResponse("Permission denied", r.URL.Path)), http.StatusForbidden)
					return
				}
				ctx := context.WithValue(r.Context(), userCtxKey, user)
				scheme := r.Header.Get("X-Forwarded-Proto")
				ctx = context.WithValue(ctx, schemeCtxKey, scheme)
				r = r.WithContext(ctx)
			}
			next.ServeHTTP(w, r)
		})
	}
}

//ForContext ..
func ForContext(ctx context.Context) *models.User {
	raw, _ := ctx.Value(userCtxKey).(*models.User)
	return raw
}

//ForSchemeContext ..
func ForSchemeContext(ctx context.Context) string {
	raw, _ := ctx.Value(schemeCtxKey).(string)
	return raw
}

//GraphReq ..
type GraphReq struct {
	OperationName string `json:"operationName"`
}

//parseBody ..
func parseBody(r *http.Request) (*GraphReq, error) {
	req := new(GraphReq)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return req, err
	}
	err = json.Unmarshal(body, req)
	if err != nil {
		return req, err
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return req, err
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//HandleIotDevice ..
func HandleIotDevice() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()
		//check accessKey and secretKey
		_, msg, err := c.ReadMessage()
		upMsg := new(pb.DeviceUpMsg)
		err = proto.Unmarshal(msg, upMsg)
		if err != nil {
			log.Println("unmarshal data error")
			return
		}
		loginMsg := upMsg.GetLoginMsg()
		id, err := deviceService.DeviceLogin(loginMsg.AccessKey, loginMsg.SecretKey)
		if id == 0 || err != nil {
			log.Println("login error:", err)
			return
		}
		device, err := deviceService.GetDeviceInfo(uint32(id))
		if err != nil {
			log.Println("get device info error:", err)
			return
		}
		deviceMsg, err := proto.Marshal(device)
		if err != nil {
			log.Println("proto marshal error:", err)
			return
		}
		err = c.WriteMessage(websocket.BinaryMessage, deviceMsg)
		if err != nil {
			log.Println("websocket write error:", err)
			return
		}
		subject := mq.SubjectDevicPrefix + strconv.FormatUint(uint64(id), 10)
		log.Println("开启监听主题：", subject)
		qsub, err := mq.GetClient().NatsConn().QueueSubscribe(subject,
			mq.GroupDevice, func(m *nats.Msg) {
				deviceMsg := new(pb.DeviceDownMSG)
				err := proto.Unmarshal(m.Data, deviceMsg)
				if err != nil {
					log.Println("unmarshal data error")
					return
				}
				switch deviceMsg.Payload.(type) {
				case *pb.DeviceDownMSG_CameraCaptureMsg:
					deviceMsg.GetCameraCaptureMsg().NatsReply = m.Reply
					requestMsg, err := proto.Marshal(deviceMsg)
					if err != nil {
						log.Println("proto marshal error:", err)
						return
					}
					err = c.WriteMessage(websocket.BinaryMessage, requestMsg)
					if err != nil {
						log.Println("websocket write error:", err)
					}
				case *pb.DeviceDownMSG_PresignedUrlReplyMsg:
					err = c.WriteMessage(websocket.BinaryMessage, m.Data)
					if err != nil {
						log.Println("websocket write error:", err)
					}
				}
			})
		if err != nil {
			log.Panicln("QueueSubscribe error:", err)
		}
		defer func() {
			err = qsub.Unsubscribe()
			if err != nil {
				log.Println("qsub2 Unsubscribe error:", err)
			}
		}()
		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			//透传到mq
			_, err = mq.GetClient().PublishAsync(mq.SubjectDeviceData, msg, utils.AckHandler)
			if err != nil {
				log.Println("send data error:", err)
			}
		}
	}
}

//CheckToken ..
func CheckToken(accessToken string) bool {
	//token验证
	if accessToken == "" {
		return false
	}
	token, err := utils.ParseJWT([]byte(consts.JWTtAccessSecret), accessToken)
	if err != nil {
		return false
	}
	data, _ := token.Claims.(*utils.MyCustomClaims)
	if data.Issuer != consts.JWTIssuer {
		return false
	}
	uid := consts.GetDecodeUID(data.UID)
	_, err = userService.GetByID(context.Background(), int64(uid))
	if err != nil {
		return false
	}
	return true
}
