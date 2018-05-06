package main

import (
	"fmt"
	"log"
	"net/http"
)

type textHandler struct {
	responseText string
}

//实现了 serveHTTP 方法就是一个 handler
func (th *textHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, th.responseText)
}

type indexHandler struct {
}

func (ih *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
func main() {
	//mux 是一个分路器
	mux := http.NewServeMux()

	//mux.handle 接受一个 pattern 和一个 handler
	mux.Handle("/", &indexHandler{})
	thWelcome := &textHandler{"Welcome to Go Web Programming"}
	mux.Handle("/welcome", thWelcome)
	thMessage := &textHandler{"net/http package is used to build web apps"}
	mux.Handle("/message", thMessage)
	log.Println("Listening...")
	http.ListenAndServe(":8000", mux)
}
