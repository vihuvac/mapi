package routers

import (
	"net/http"

	"../controllers"
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
	router.Use(accessControlMiddleware)
	router = SetRoutes(router)
	return router
}

func accessControlMiddleware(server http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		server.ServeHTTP(w, r)
	})
}
