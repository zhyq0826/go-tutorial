package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func (m *HelloHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello")
}

type WorldHandler struct{}

func (m *WorldHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "world")
}

func main() {
	h := HelloHandler{}
	w := WorldHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/hello", &h)
	http.Handle("/world", &w)

	server.ListenAndServe()
}
