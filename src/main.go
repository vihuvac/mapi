package main

import (
	"log"
	"net/http"
	"os"

	"./databases/postgres"
	"./routers"
)

func main() {
	port := os.Getenv("PORT")

	_, err := postgres.Init()
	if err != nil {
		log.Println("The database connection failed! Aborting...")
		log.Fatal(err)
	}

	log.Println("Database connection established.")

	router := routers.InitRoutes()
	http.ListenAndServe(":"+port, router)
}
