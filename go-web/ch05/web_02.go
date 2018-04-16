package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func ifHandler(rw http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("if.html"))
	rand.Seed(time.Now().Unix())
	t.Execute(rw, rand.Intn(10) > 5)
}

func rangeHandler(rw http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("range.html"))
	t.Execute(rw, []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"})
}

func setHandler(rw http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("set.html"))
	t.Execute(rw, "HELLO")
}

func includeHandler(rw http.ResponseWriter, req *http.Request) {
	t := template.Must(template.ParseFiles("t1.html", "t2.html"))
	t.Execute(rw, "hello world")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}

	http.HandleFunc("/if", ifHandler)
	http.HandleFunc("/range", rangeHandler)
	http.HandleFunc("/set", setHandler)
	http.HandleFunc("/include", includeHandler)
	server.ListenAndServe()
}
