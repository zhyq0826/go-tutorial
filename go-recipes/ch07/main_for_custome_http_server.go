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
	//通过 http.Server 可以定制 http server
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
	log.Println("Listening...")
	// listenAndServe 内部使用默认的 serveMux
	server.ListenAndServe()
}
