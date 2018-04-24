package main

import (
	"html/template"
	"net/http"
	// "time"
)

func contextHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("context_aware.html")
	content := `I asked: <i>"What's up?"</i>`
	t.Execute(w, content)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}

	http.HandleFunc("/context", contextHandler)
	server.ListenAndServe()
}
