package main

import (
	"log"
	"net/http"

	"github.com/zhyq0826/go-tutorial/go-recipes/ch07-2/common"
	"github.com/zhyq0826/go-tutorial/go-recipes/ch07-2/routers"
)

// Entry point of the program
func main() {

	// Calls startup logic
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()
	// Create the Server
	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}
	log.Println("Listening...")
	// Running the HTTP Server
	server.ListenAndServe()
}
