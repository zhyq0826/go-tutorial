package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"

	"github.com/zhyq0826/go-tutorial/devops/apps"
	_ "github.com/zhyq0826/go-tutorial/devops/db"
)

func main() {
	router := apps.InitRouter()
	handler := handlers.LoggingHandler(os.Stdout, router)
	server := &http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: handler,
	}
	log.Println("Listening...")
	// start run http server
	server.ListenAndServe()
}
