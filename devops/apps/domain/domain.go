package domain

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	helper "github.com/zhyq0826/go-tutorial/devops/helpers/devops"
)

// InitRoutes init domain app route
func InitRoutes(router *mux.Router) *mux.Router {
	appRouter := router.PathPrefix("/domain").Subrouter()
	appRouter.HandleFunc("/list", GetDomains)
	return router
}

// GetDomains domain list handle
func GetDomains(w http.ResponseWriter, r *http.Request) {
	domains := helper.QueryDomain(0, 10, "", "")
	jsonData, _ := json.Marshal(domains)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
