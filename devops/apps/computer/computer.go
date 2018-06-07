package computer

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
	appRouter := router.PathPrefix("/computer").Subrouter()
	appRouter.HandleFunc("/list", getComputers).Methods("GET")
	appRouter.HandleFunc("/create", createComputer).Methods("POST")
	appRouter.HandleFunc("/{id:[0-9]+}", updateComputer).Methods("PUT")
	appRouter.HandleFunc("/{id:[0-9]+}", deleteComputer).Methods("DELETE")
	return router
}

// getComputers computer list handle
func getComputers(w http.ResponseWriter, r *http.Request) {
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

	ret := helper.QueryComputer(skip, limit)
	jsonData, _ := json.Marshal(ret)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

// createComputer handle
func createComputer(w http.ResponseWriter, r *http.Request) {
	dataSource := resources.ComputerForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		log.Println(err)
	}
	helper.CreateComputer(dataSource)
}

// updateComputer handle
func updateComputer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dataSource := resources.ComputerForm{}
	err := json.NewDecoder(r.Body).Decode(&dataSource)
	if err != nil {
		log.Println(err)
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("convert id to fails")
	}
	helper.UpdateComputer(id, dataSource)
}

// deleteComputer handle
func deleteComputer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("convert id to fails")
	}
	helper.DeleteComputer(id)
}
