package handler

import (
	dblayer "filespan/db"
	"filespan/util"
	"io/ioutil"
	"net/http"
)

const (
	pwd_salt = "#890"
)

func SignupHandler(w http.ResponseWriter,r *http.Request) {
	// 判断请求方法
	if r.Method == http.MethodGet {
		data,err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}
	r.ParseForm()
	username := r.Form.Get("username")
	passwd := r.Form.Get("password")
	if len(username) < 3 || len(passwd) < 5 {
		w.Write([]byte("Invalied parameter"))
		return
	}
	// 加密
	enc_passwd := util.Sha1([]byte(passwd+pwd_salt))

	ret := dblayer.UserSign(username,enc_passwd)
	if ret {
		w.Write([]byte("SUCCESS"))
	} else {
		w.Write([]byte("Failed sign up "))
	}



}