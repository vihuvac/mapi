package main

import (
	"log"
	"os"

	"./controllers"
	"./databases"
	"./databases/postgres"
	"./middlewares"
	"./routers"
	"./services"
)

var (
	characterStorage    databases.CharacterStorage      = databases.NewPostgresQLStorage()
	characterService    services.CharacterService       = services.NewCharacterService(characterStorage)
	characterController controllers.CharacterController = controllers.NewCharacterController(characterService)
	router              routers.Router                  = routers.NewMuxRouter()
)

func main() {
	port := os.Getenv("PORT")

	defer postgres.DB.Close()

	_, err := postgres.Init()
	if err != nil {
		log.Println("The database connection failed! Aborting...")
		log.Fatal(err)
	}
	log.Println("Database connection established.")

	// Apply middlewares to all routes.
	router.Use(middlewares.AccessControl)

	// Enable routes for characters.
	router.GET("/api/characters", characterController.GetCharacters)
	router.POST("/api/characters", middlewares.IsAuthorized(characterController.CreateCharacter))

	// Run the mux server.
	router.Serve(port)
}
