package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func HanderUpload() {
	http.HandleFunc("/upload", upload)
}

func upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		strtime := strconv.FormatInt(crutime, 10)
		println("strtime:", strtime)
		//io.WriteString(w, strtime)
		token := fmt.Sprintf("%x", h.Sum(nil))
		println(token + "------------------------------------")
		t, _ := template.ParseFiles("resources/upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)

		file, header, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		//关闭multipart.File
		defer file.Close()
		fmt.Fprintf(w, "%v", header.Header)
		//打开下载文件，如果不存在则创建
		f, err := os.OpenFile(
			"./upload/"+header.Filename,
			os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		//关闭os.File
		defer f.Close()
		io.Copy(f, file)
	}
}
