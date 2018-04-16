package main

import (
	"html/template"
	"net/http"
	"time"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func funcHandler(rw http.ResponseWriter, req *http.Request) {
	funcMap := template.FuncMap{
		"fdate": formatDate,
	}
	t := template.New("func.html").Funcs(funcMap)
	t, _ = t.ParseFiles("func.html")
	t.Execute(rw, time.Now())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}

	http.HandleFunc("/func", funcHandler)
	server.ListenAndServe()
}
