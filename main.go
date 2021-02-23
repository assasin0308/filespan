package main

import (
	"filespan/handler"
	"fmt"
	"net/http"
)

func main() {
	// 路由规则配置
	http.HandleFunc("/file/upload",handler.UploadHandler)
	http.HandleFunc("/file/upload/success",handler.UploadSuccessHandler)
	http.HandleFunc("/file/meta",handler.GetFileMetaHandler)
	// 端口监听
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Printf("Failed to start server,error:%s",err.Error())
	}
}
