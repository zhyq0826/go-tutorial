package consul

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/hashicorp/consul/api"

	"github.com/gorilla/mux"
)

// InitRoutes init domain app route
func InitRoutes(router *mux.Router) *mux.Router {
	appRouter := router.PathPrefix("/consul").Subrouter()
	appRouter.HandleFunc("/service/all", getAllService).Methods("GET")
	appRouter.HandleFunc("/service/list", getService).Methods("GET")
	return router
}

var consulClient, err = api.NewClient(api.DefaultConfig())

func getAllService(w http.ResponseWriter, r *http.Request) {
	catalog := consulClient.Catalog()
	services, _, err := catalog.Services(nil)
	if err != nil {
		log.Println(err)
	}

	jsonData, _ := json.Marshal(services)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func getService(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	tag := r.URL.Query().Get("tag")
	if name == "" && tag == "" {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{}`))
		return
	}

	catalog := consulClient.Catalog()
	services, _, err := catalog.Service(name, tag, nil)
	if err != nil {
		log.Println(err)
	}

	list := make([]map[string]interface{}, 0)
	for _, s := range services {
		ret := make(map[string]interface{})
		ret["service_id"] = s.ServiceID
		ret["service_name"] = s.ServiceName
		ret["service_address"] = s.ServiceAddress
		ret["service_port"] = s.ServicePort
		ret["service_tags"] = s.ServiceTags
		ret["datacenter"] = s.Datacenter
		list = append(list, ret)
	}

	jsonData, _ := json.Marshal(list)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
