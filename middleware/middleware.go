package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/9d77v/pdc/iot/sdk/pb"
	"github.com/9d77v/pdc/models"
	"github.com/9d77v/pdc/models/nats"
	"github.com/9d77v/pdc/services"
	"github.com/9d77v/pdc/utils"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

var userCtxKey = &contextKey{"user"}
var schemeCtxKey = &contextKey{"scheme"}

type contextKey struct {
	name string
}

//GraphReq ..
type GraphReq struct {
	OperationName string `json:"operationName"`
}

var (
	userService   = services.UserService{}
	deviceService = services.DeviceService{}
)

var publicOperationArr = []string{"login", "refreshToken", "IntrospectionQuery"}
var permissonMap = map[string][]int{
	"presignedUrl":                   {models.RoleAdmin, models.RoleManager, models.RoleUser},
	"users":                          {models.RoleAdmin},
	"userInfo":                       {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
	"videos":                         {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
	"videoSerieses":                  {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
	"searchVideo":                    {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
	"things":                         {models.RoleAdmin, models.RoleManager, models.RoleUser},
	"thingSeries":                    {models.RoleAdmin, models.RoleManager, models.RoleUser},
	"thingAnalyze":                   {models.RoleAdmin, models.RoleManager, models.RoleUser},
	"createUser":                     {models.RoleAdmin},
	"updateUser":                     {models.RoleAdmin},
	"updateProfile":                  {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
	"updatePassword":                 {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
	"createVideo":                    {models.RoleAdmin, models.RoleManager},
	"updateVideo":                    {models.RoleAdmin, models.RoleManager},
	"createEpisode":                  {models.RoleAdmin, models.RoleManager},
	"updateEpisode":                  {models.RoleAdmin, models.RoleManager},
	"updateSubtitle":                 {models.RoleAdmin, models.RoleManager},
	"updateMobileVideo":              {models.RoleAdmin, models.RoleManager},
	"createVideoSeries":              {models.RoleAdmin, models.RoleManager},
	"updateVideoSeries":              {models.RoleAdmin, models.RoleManager},
	"createVideoSeriesItem":          {models.RoleAdmin, models.RoleManager},
	"updateVideoSeriesItem":          {models.RoleAdmin, models.RoleManager},
	"createThing":                    {models.RoleAdmin, models.RoleManager, models.RoleUser},
	"updateThing":                    {models.RoleAdmin, models.RoleManager, models.RoleUser},
	"recordHistory":                  {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
	"historyInfo":                    {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
	"histories":                      {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
	"createDeviceModel":              {models.RoleAdmin, models.RoleManager},
	"updateDeviceModel":              {models.RoleAdmin, models.RoleManager},
	"createAttributeModel":           {models.RoleAdmin, models.RoleManager},
	"updateAttributeModel":           {models.RoleAdmin, models.RoleManager},
	"deleteAttributeModel":           {models.RoleAdmin, models.RoleManager},
	"createTelemetryModel":           {models.RoleAdmin, models.RoleManager},
	"updateTelemetryModel":           {models.RoleAdmin, models.RoleManager},
	"deleteTelemetryModel":           {models.RoleAdmin, models.RoleManager},
	"deviceModels":                   {models.RoleAdmin, models.RoleManager},
	"createDevice":                   {models.RoleAdmin, models.RoleManager},
	"updateDevice":                   {models.RoleAdmin, models.RoleManager},
	"devices":                        {models.RoleAdmin, models.RoleManager},
	"createDeviceDashboard":          {models.RoleAdmin, models.RoleManager},
	"updateDeviceDashboard":          {models.RoleAdmin, models.RoleManager},
	"deleteDeviceDashboard":          {models.RoleAdmin, models.RoleManager},
	"addDeviceDashboardTelemetry":    {models.RoleAdmin, models.RoleManager},
	"removeDeviceDashboardTelemetry": {models.RoleAdmin, models.RoleManager},
	"deviceDashboards":               {models.RoleAdmin, models.RoleManager},
}

//Auth ..
func Auth() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/api" {
				req, err := parseBody(r)
				if err != nil {
					http.Error(w, "Invalid request", http.StatusBadRequest)
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
					http.Error(w, "Invalid token", http.StatusUnauthorized)
					return
				}
				accessToken = strings.TrimPrefix(accessToken, "Bearer ")
				token, err := utils.ParseJWT([]byte(models.JWTtAccessSecret), accessToken)
				if err != nil {
					http.Error(w, "Invalid token", http.StatusUnauthorized)
					return
				}
				data, _ := token.Claims.(*utils.MyCustomClaims)
				if data.Issuer != models.JWTIssuer {
					http.Error(w, "Invalid token", http.StatusUnauthorized)
					return
				}

				user, err := userService.GetByID(r.Context(), data.UID)
				if err != nil {
					http.Error(w, "Invalid token", http.StatusUnauthorized)
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
					http.Error(w, "Permission denied", http.StatusForbidden)
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
		loginMsg := new(pb.LoginMSG)
		err = proto.Unmarshal(msg, loginMsg)
		if err != nil {
			log.Println("unmarshal data error")
			return
		}
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

		for {
			_, msg, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			//透传到mq
			_, err = nats.Client.PublishAsync(nats.SubjectDeviceData, msg, utils.AckHandler)
			if err != nil {
				log.Println("send data error:", err)
			}
		}
	}
}
