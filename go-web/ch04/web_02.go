package main

import (
	// "io/ioutil"
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func writeHandler(rw http.ResponseWriter, req *http.Request) {
	str := `
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <title>hello go</title>
        </head>
        <body>
            hello go
        </body>
        </html>
    `
	rw.Write([]byte(str))
}

func writeHeaderHandler(rw http.ResponseWriter, _ *http.Request) {
	rw.WriteHeader(501)
	fmt.Fprintln(rw, "No such service, try next door")
}

func headerHandler(rw http.ResponseWriter, _ *http.Request) {
	rw.Header().Set("Location", "https://baidu.com")
	rw.WriteHeader(301)
}

func jsonHandler(rw http.ResponseWriter, _ *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "shehong",
		Threads: []string{"first", "second", "third"},
	}

	d, _ := json.Marshal(post)
	rw.Write(d)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}

	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/writeHeader", writeHeaderHandler)
	http.HandleFunc("/header", headerHandler)
	http.HandleFunc("/json", jsonHandler)
	server.ListenAndServe()
}
