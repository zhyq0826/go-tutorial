package routers

import (
	"github.com/zhyq0826/go-tutorial/go-recipes/ch07-2/controllers"

	"github.com/gorilla/mux"
)

// SetUserRoutes registers routes for user entity
func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users", controllers.Register).Methods("POST")
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
	return router
}
