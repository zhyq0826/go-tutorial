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
	mux.Handle("/", http.HandlerFunc(index))
	mux.Handle("/welcome", http.HandlerFunc(welcome))
	mux.Handle("/message", http.HandlerFunc(message))
	log.Println("Listening...")
	http.ListenAndServe(":8000", mux)
}
