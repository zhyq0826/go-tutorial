package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

func (m *MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello world")
}

func main() {
	handler := MyHandler{}

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler, //ServeHTTP 的 method receiver 是 pointer
		// Handler 默认是一个 defaultServeMux instance，现在换成了支持 MyHandler
		// 这样所有的请求都交给了 myhandler 处理，即 server 中默认的 Handler
	}

	server.ListenAndServe()
}
