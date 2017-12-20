package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func FormPage() {
	http.HandleFunc("/", homeHandle)
	http.HandleFunc("/login", loginHandle)
}

func homeHandle(w http.ResponseWriter, r *http.Request) {
	//输出到客户端
	fmt.Fprintf(w, "欢迎来到首页")
}

func loginHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method", r.Method)
	if r.Method == "GET" {
		//获取表单
		curtime := time.Now().Unix() //获取当前时间
		h := md5.New()               //生成MD5
		io.WriteString(w, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		println("token:", token, "----------------------")
		t, _ := template.ParseFiles("resources/login.gtpl")
		t.Execute(w, token) //预防表单多次提交
		log.Println(t.Execute(w, nil))
	} else {
		//提交表单
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			//验证token的合法性
		} else {
			//不存在token报错
		}
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
