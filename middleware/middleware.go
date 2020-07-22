package middleware

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/9d77v/pdc/models"
	"github.com/9d77v/pdc/services"
	"github.com/9d77v/pdc/utils"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

//GraphReq ..
type GraphReq struct {
	OperationName string `json:"operationName"`
}

var userService = services.UserService{}

var publicOperationArr = []string{"login", "refreshToken"}
var permissonMap = map[string][]int{
	"presignedUrl":   {models.RoleAdmin, models.RoleManager, models.RoleUser},
	"users":          {models.RoleAdmin},
	"userInfo":       {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
	"videos":         {models.RoleAdmin, models.RoleManager, models.RoleUser, models.RoleGuest},
	"things":         {models.RoleAdmin, models.RoleManager, models.RoleUser},
	"thingSeries":    {models.RoleAdmin, models.RoleManager, models.RoleUser},
	"thingAnalyze":   {models.RoleAdmin, models.RoleManager, models.RoleUser},
	"createUser":     {models.RoleAdmin},
	"updateUser":     {models.RoleAdmin},
	"createVideo":    {models.RoleAdmin, models.RoleManager},
	"updateVideo":    {models.RoleAdmin, models.RoleManager},
	"createEpisode":  {models.RoleAdmin, models.RoleManager},
	"updateEpisode":  {models.RoleAdmin, models.RoleManager},
	"updateSubtitle": {models.RoleAdmin, models.RoleManager},
	"createThing":    {models.RoleAdmin, models.RoleManager, models.RoleUser},
	"updateThing":    {models.RoleAdmin, models.RoleManager, models.RoleUser},
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
