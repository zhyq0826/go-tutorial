package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	helper "github.com/zhyq0826/go-tutorial/devops/helpers/devops"
	"github.com/zhyq0826/go-tutorial/devops/resources"
)

// InitRoutes init domain app route
func InitRoutes(router *mux.Router) *mux.Router {
	appRouter := router.PathPrefix("/service").Subrouter()
	appRouter.HandleFunc("/list", getServices).Methods("GET")
	appRouter.HandleFunc("/create", createService).Methods("POST")
	appRouter.HandleFunc("/{id:[0-9]+}", updateService).Methods("PUT")
	appRouter.HandleFunc("/{id:[0-9]+}", deleteService).Methods("DELETE")
	return router
}

// getServices service list handle
func getServices(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var skip, limit = 0, 10
	var err error
	if q.Get("skip") != "" {
		skip, err = strconv.Atoi(q.Get("skip"))
		if err != nil {
			log.Println("skip parse error")
		}
	}

	if q.Get("limit") != "" {
		limit, err = strconv.Atoi(q.Get("limit"))
		if err != nil {
			log.Println("limit parse error")
		}
	}

	ret := helper.QueryService(skip, limit, q.Get("name"), q.Get("url"))
	jsonData, _ := json.Marshal(ret)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// createService handle
func createService(w http.ResponseWriter, r *http.Request) {
	dataSource := resources.ServiceForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		log.Println(err)
	}
	helper.CreateService(dataSource)
}

// updateService handle
func updateService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dataSource := resources.ServiceForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		log.Println(err)
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("convert id to fails")
	}
	helper.UpdateService(id, dataSource)
}

// deleteService handle
func deleteService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("convert id to fails")
	}
	helper.DeleteService(id)
}
