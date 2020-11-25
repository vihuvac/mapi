package main

import (
	"net/http"
	"os"

	"./routers"
)

func main() {
	port := os.Getenv("PORT")
	router := routers.InitRoutes()
	http.ListenAndServe(":"+port, router)
}
