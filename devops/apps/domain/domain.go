package domain

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
	appRouter := router.PathPrefix("/domain").Subrouter()
	appRouter.HandleFunc("/list", getDomains)
	appRouter.HandleFunc("/create", createDomain)
	appRouter.HandleFunc("/{id:[0-9]+}", updateDomain)
	return router
}

// getDomains domain list handle
func getDomains(w http.ResponseWriter, r *http.Request) {
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

	domains := helper.QueryDomain(skip, limit, q.Get("name"), q.Get("url"))
	jsonData, _ := json.Marshal(domains)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// createDomain handle
func createDomain(w http.ResponseWriter, r *http.Request) {
	dataSource := resources.DomainForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		log.Println(err)
	}
	helper.CreateDomain(dataSource.Name, dataSource.URL, dataSource.Private)
}

// updateDomain handle
func updateDomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dataSource := resources.DomainForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		log.Println(err)
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("convert id to fails")
	}
	helper.UpdateDomain(id, dataSource.Name, dataSource.URL, dataSource.Private)
}
