package main

import (
	"log"
	"net/http"

	"github.com/zhyq0826/go-tutorial/devops/apps"
	_ "github.com/zhyq0826/go-tutorial/devops/db"
)

func main() {
	router := apps.InitRouter()
	server := &http.Server{
		Addr:    "127.0.0.1:8300",
		Handler: router,
	}

	log.Println("Listening...")
	// start run http server
	server.ListenAndServe()
}
