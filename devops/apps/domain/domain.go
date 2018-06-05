package domain

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// InitRoutes init domain app route
func InitRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/domain/list", GetDomains)
	// appRouter := mux.NewRouter()
	// appRouter.HandleFunc("/list", GetDomains)
	// router.PathPrefix("/domain").Handler(handlers.LoggingHandler(os.Stdout, appRouter))
	return router
}

// GetDomains domain list handle
func GetDomains(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{"value": "hello world"}
	jsonData, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
