package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello world %s", request.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)
	// 第二个参数是一个 handler，默认一般一个 serverMux
	http.ListenAndServe(":8080", nil)
}
