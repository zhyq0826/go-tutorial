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
	appRouter := router.PathPrefix("/v1/domain").Subrouter()
	appRouter.HandleFunc("/list", getDomains)
	appRouter.HandleFunc("/create", createDomain).Methods("POST")
	appRouter.HandleFunc("/{id:[0-9]+}", updateDomain).Methods("PUT")
	appRouter.HandleFunc("/{id:[0-9]+}", deleteDomain).Methods("DELETE")
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
	// var dataSource []resources.DomainForm
	// for _, v := range domains {
	// 	vv := v
	// 	dest := resources.DomainForm{}
	// 	utils.CopyStruct(&vv, &dest)
	// 	dataSource = append(dataSource, dest)
	// }
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
	helper.CreateDomain(dataSource.Name, dataSource.Host, dataSource.Private, dataSource.IP)
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
	helper.UpdateDomain(id, dataSource.Name, dataSource.Host, dataSource.Private)
}

// deleteDomain handle
func deleteDomain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("convert id to fails")
	}
	helper.DeleteDomain(id)
}
