package routers

import (
	"../controllers"
	"../middlewares"
	"github.com/gorilla/mux"
)

// SetRoutes sets all the available routes.
func SetRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/api/characters", controllers.GetCharacters).Methods("GET")
	return router
}

// InitRoutes initializes the routes.
func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router.Use(middlewares.AccessControl)
	router = SetRoutes(router)
	return router
}
