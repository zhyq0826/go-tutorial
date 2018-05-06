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
	mux := http.NewServeMux()
	// handleFunc 是一个 adapter 可以把 一个普通 func 改造成一个 handler
	// serveMux 的 handleFunc 可以扮演的角色类似 handleFunc，不过他还可以完成 handler 到 serveMux 的注册
	mux.HandleFunc("/", index)
	mux.HandleFunc("/welcome", welcome)
	mux.HandleFunc("/message", message)
	log.Println("Listening...")
	http.ListenAndServe(":8000", mux)
}
