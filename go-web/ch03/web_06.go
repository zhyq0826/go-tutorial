package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello")
}

// wrap a handler 并且返回一个 Handler
func log(h http.Handler) http.Handler {
	// 进行转换
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("handler func called - ", name)
		h.ServeHTTP(writer, request)
	})
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/hello", log(&HelloHandler{}))
	server.ListenAndServe()
}
