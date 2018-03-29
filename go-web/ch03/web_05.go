package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}

func hello(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "hello")
}

type WorldHandler struct{}

func world(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "world")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}
