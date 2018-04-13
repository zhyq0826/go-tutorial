package main

import (
	"fmt"
	"io/ioutil"
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
	fmt.Fprintln(rw, req.MultipartForm) //支持带有文件的表单, 不包括 url query
}

func valueHandler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, req.FormValue("first_name"))
	fmt.Fprintln(rw, req.PostFormValue("first_name"))
	fmt.Fprintln(rw, req.Form)          // 包括 url query
	fmt.Fprintln(rw, req.PostForm)      // 只包括 form
	fmt.Fprintln(rw, req.MultipartForm) // form value 执行之前是调用 paraseMutilForm
}

func fileHandler(rw http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(1014)
	fileheader := req.MultipartForm.File["myfile"][0]
	file, err := fileheader.Open()
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(rw, string(data))
		}
	}
}

func fileHandler2(rw http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("myfile")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(rw, string(data))
		}
	}
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
	http.HandleFunc("/file", fileHandler)
	http.HandleFunc("/file2", fileHandler2)
	error := server.ListenAndServe()
	fmt.Println(error)
}
