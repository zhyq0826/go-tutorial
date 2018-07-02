package app

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
	appRouter := router.PathPrefix("/v1/app").Subrouter()
	appRouter.HandleFunc("/list", getApps).Methods("GET")
	appRouter.HandleFunc("/create", createApp).Methods("POST")
	appRouter.HandleFunc("/{id:[0-9]+}", oneApp).Methods("GET")
	appRouter.HandleFunc("/{id:[0-9]+}", updateApp).Methods("PUT")
	appRouter.HandleFunc("/{id:[0-9]+}", deleteApp).Methods("DELETE")
	return router
}

// getApps service list handle
func getApps(w http.ResponseWriter, r *http.Request) {
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

	ret := helper.QueryApp(skip, limit, q.Get("name"), q.Get("url"))
	jsonData, _ := json.Marshal(ret)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// createApp handle
func createApp(w http.ResponseWriter, r *http.Request) {
	dataSource := resources.AppForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		log.Println(err)
	}
	helper.CreateApp(dataSource)
}

// updateApp handle
func updateApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dataSource := resources.AppForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		log.Println(err)
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("convert id to fails")
	}
	helper.UpdateApp(id, dataSource)
}

func oneApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("convert id to fails")
	}
	ret := helper.QueryOneApp(id)
	jsonData, _ := json.Marshal(ret)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// deleteApp handle
func deleteApp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("convert id to fails")
	}
	helper.DeleteApp(id)
}
