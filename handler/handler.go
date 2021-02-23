package handler

import (
	"filespan/meta"
	"filespan/util"
	"fmt"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func UploadHandler(w http.ResponseWriter,r *http.Request) {
	if r.Method == "GET" {
		// 返回上传的HTML页面
		data,err := ioutil.ReadFile("./static/view/upload.html")
		if err != nil {
			io.WriteString(w,"Internal Server Err")
			return
		}
		io.WriteString(w,string(data))

	}else if r.Method == "POST" {
		// 接收文件流存储到本地目录
		file,head,err := r.FormFile("file")
		if err != nil {
			fmt.Printf("Failed to get data,error:%s",err.Error())
			return
		}
		defer file.Close()
		fileMeta := meta.FileMeta{
			FileName:head.Filename,
			Location:`E:\go-project\src\filespan\uploaded\`,
			UploadTime:time.Now().Format("2006-01-02 15:04:05"),
		}

		// 创建新文件
		newfile,err := os.Create(fileMeta.Location + head.Filename)
		if err != nil {
			fmt.Printf("Failed to create file,error:%s",err.Error())
			return
		}
		defer newfile.Close()


		// 拷贝文件内容至新文件
		fileMeta.FileSize,err = io.Copy(newfile,file)
		if err != nil {
			fmt.Printf("Failed to save data into file,error:%s",err.Error())
			return
		}
		// 计算文件的hash值
		newfile.Seek(0,0)
		fileMeta.FileSha1 = util.FileSha1(newfile)
		meta.UpdateFileMeta(fileMeta)

		// 上传结束,重定向
		http.Redirect(w,r,"/file/upload/success",http.StatusFound)
	}
}

func UploadSuccessHandler(w http.ResponseWriter,r *http.Request){
	io.WriteString(w,"Upload finished! ")
}

// 获取单个文件的元信息
func GetFileMetaHandler(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	filehash := r.Form["filehash"][0]
	fileMeta := meta.GetFileMeta(filehash)
	data,err := json.Marshal(fileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// 文件下载
func DownLoadHandler(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	filehash := r.Form.Get("filehash")
	fileMeta := meta.GetFileMeta(filehash)
	f,err := os.Open(fileMeta.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer  f.Close()
	data,err := ioutil.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/octect-stream")
	w.Header().Set("Content-Description","attachment;filename=\""+fileMeta.FileName + "\"")
	w.Write(data)
}


// 更新元信息(重命名)
func FileMetaUpdateHandler(w http.ResponseWriter,r *http.Request){
	r.ParseForm()
	opType := r.Form.Get("op")
	fileSha1 := r.Form.Get("filehash")
	newFileName := r.Form.Get("filename")
	if opType != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	currentFileMeta := meta.GetFileMeta(fileSha1)
	currentFileMeta.FileName = newFileName
	meta.UpdateFileMeta(currentFileMeta)


	data,err := json.Marshal(currentFileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// 文件元信息删除
func FileDeletehandler(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	fileSha1 := r.Form.Get("filehash")

	fMeta := meta.GetFileMeta(fileSha1)
	os.Remove(fMeta.Location)

	meta.RemoveFileMeta(fileSha1)

	w.WriteHeader(http.StatusOK)

}