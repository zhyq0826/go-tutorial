package main

import (
	"fmt"
	"net/http"
)

func headerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Header)

	fmt.Fprintln(w, r.Header.Get("Accept-Encoding")) //gzip, deflate, br
	fmt.Fprintln(w, r.Header["Accept-Encoding"])     //[gzip, deflate, br]
}

func bodyHandler(rw http.ResponseWriter, req *http.Request) {
	length := req.ContentLength
	body := make([]byte, length)
	req.Body.Read(body)
	fmt.Fprintln(rw, string(body))
}

func formHandler(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Fprintln(rw, req.Form)     //包含 url 中的 query 参数
	fmt.Fprintln(rw, req.PostForm) //只包含 post form 的数据, post form 只支持 x-www-form-urlencode 的 form 表单
}

func multiFormHandler(rw http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(1024)
	fmt.Fprintln(rw, req.MultipartForm) //支持带有文件的表单
}

func valueHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, req.FormValue("first_name"))
	fmt.Fprintln(rw, req.PostFormValue("first_name"))
	fmt.Fprintln(rw, req.Form)
	fmt.Fprintln(rw, req.PostForm)
	fmt.Fprintln(rw, req.MultipartForm)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}

	http.HandleFunc("/header", headerHandler)
	http.HandleFunc("/body", bodyHandler)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/multiform", multiFormHandler)
	http.HandleFunc("/value", valueHandler)
	error := server.ListenAndServe()
	fmt.Println(error)
}
