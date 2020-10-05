package middleware

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//User public user info
type User struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
}

//HandleCard get public user info
func HandleCard() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		user := new(User)
		arr := strings.Split(req.RequestURI, "/")
		if len(arr) != 3 {
			io.WriteString(w, "request error")
			return
		}
		id, err := strconv.ParseInt(arr[2], 10, 64)
		if err != nil {
			log.Println("conver id to int error:", err)
		}
		u, err := userService.GetByID(context.Background(), id)
		if err != nil {
			log.Println("get user error:", err)
		}
		user.Avatar = u.Avatar
		user.Name = u.Name
		ret, err := json.Marshal(user)
		if err != nil {
			log.Println("json marshal faield:", err)
		}
		w.Header().Set("Content-type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-type")
		io.WriteString(w, string(ret))
	}
}
