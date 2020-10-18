package middleware

import (
	"context"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/9d77v/pdc/dtos"
)

//User public user info
type User struct {
	UID    uint   `json:"uid"`
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
	Gender int    `json:"gender"`
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
		user.UID = u.ID
		scheme := req.Header.Get("X-Forwarded-Proto")
		user.Avatar = dtos.GetOSSPrefix(scheme) + u.Avatar
		user.Name = u.Name
		user.Gender = u.Gender
		ret, err := json.Marshal(user)
		if err != nil {
			log.Println("json marshal faield:", err)
		}
		accept := req.Header.Get("Accept")
		if strings.Contains(accept, "application/json") {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-type")
			w.Header().Set("Content-type", "application/json")
			io.WriteString(w, string(ret))
			return
		}
		if req.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		w.Header().Set("Content-type", "text/html")
		tem, err := template.ParseFiles("tpls/user.html")
		if err != nil {
			log.Println("读取文件失败,err", err)
			return
		}
		tem.Execute(w, user)
	}
}
