package main

import (
	"html/template"
	"net/http"
)

func templateHandler(rw http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(rw, "hello world")
}

func templateHandler2(rw http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("tmpl.html"))
	t.Execute(rw, "hello world")
}

func templateHandler3(rw http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("tmpl.html", "tmpl2.html"))
	t.ExecuteTemplate(rw, "tmpl2.html", "hello world")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}

	http.HandleFunc("/template1", templateHandler)
	http.HandleFunc("/template2", templateHandler2)
	http.HandleFunc("/template3", templateHandler3)
	server.ListenAndServe()
}
