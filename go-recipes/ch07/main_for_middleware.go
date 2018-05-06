package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// HTTP middlewares are pluggable and self-contained piece of code that wrap HTTP handlers of web
// applications. These are like typical HTTP handlers, but they wrap another HTTP handler, typically normal
// application handlers to provide shared behaviors to web applications. It works as an another layer in
// the HTTP request-handling cycle to inject some pluggable code for executing shared behaviors such as
// authentication and authorization, logging, caching, and so on.
//  Here is the basic pattern for writing HTTP middleware

// 中间件的职责就是包装 http.Handler 在不侵入代码的前提下实现特定的功能，go 的 middlware 是接受 handler 再返回 handler

// func middlewareHandler(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Middleware logic goes here before executing application handler
// 		next.ServeHTTP(w, r)
// 		// Middleware logic goes here after executing application handler
// 	})
// }

// loggingHandler is an HTTP Middleware that logs HTTP requests.
func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Middleware logic before executing given Handler
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		// Middleware logic after executing given Handler
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

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
	http.Handle("/", loggingHandler(http.HandlerFunc(index)))
	http.Handle("/welcome", loggingHandler(http.HandlerFunc(welcome)))
	http.Handle("/message", loggingHandler(http.HandlerFunc(message)))
	log.Println("Listening...")
	http.ListenAndServe(":8000", nil)
}
