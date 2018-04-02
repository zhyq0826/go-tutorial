package main

import (
	"fmt"
	"net/http"
)

func headerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.Header)
}

func bodyHandler(rw http.ResponseWriter, req *http.Request) {
	length := req.ContentLength
	body := make([]byte, length)
	req.Body.Read(body)
	fmt.Fprintln(rw, string(body))
}

func formHandler(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Fprintln(rw, req.Form)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}

	http.HandleFunc("/header", headerHandler)
	http.HandleFunc("/body", bodyHandler)
	http.HandleFunc("/form", formHandler)
	error := server.ListenAndServe()
	fmt.Println(error)
}
