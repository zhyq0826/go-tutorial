package apps

import (
	"github.com/gorilla/mux"
	"github.com/zhyq0826/go-tutorial/devops/apps/domain"
)

// InitRouter init all app route
func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router = domain.InitRoutes(router)
	return router
}
