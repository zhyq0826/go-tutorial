package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(
		"Content-Type",
		"text/html",
	)
	html :=
		`<doctype html>
 <html>
 <head>
 <title>Hello Gopher</title>
 </head>
 <body>
 <b>Hello Gopher!</b>
 <p>
 <a href="/welcome">Welcome</a> | <a href="/message">Message</a>
 </p>
 </body>
</html>`
	fmt.Fprintf(w, html)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go Web Programming")
}
func message(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "net/http package is used to build web apps")
}

func main() {
	// http 提供了一个默认的 mux 叫 DefaultServeMux，这是一个 ServeMux 的实例
	// 使用 DefaultServeMux 就不需要在代码中声明 ServeMux 了
	http.HandleFunc("/", index)
	http.HandleFunc("/welcome", welcome)
	http.HandleFunc("/message", message)
	log.Println("Listening...")
	// If you provide a nil value, package http will take DefaultServeMux as the ServeMux value. When you
	// work with DefaultServeMux as the ServeMux value, you can use the function http.HandleFunc to register
	// a handler function for the given URL pattern. Inside the function http.HandleFunc , it calls the function
	// HandleFunc of DefaultServeMux . The HandleFunc of ServeMux then calls the function Handle of ServeMux by
	// providing the http.HandlerFunc call using the given handler function

	// nil 表示 http 包将使用默认的 defaultServeMux，使用 http.handleFunc 将会注册 handler 到默认的 serveMux
	http.ListenAndServe(":8000", nil)
}
