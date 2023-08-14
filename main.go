package main

import (
	"log"
	"mylovepp/repositories"
	"mylovepp/routes"
)

func main() {
	// setup database connection
	client, err := repositories.InitialDbConnection()
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	} else {
		defer client.Close()

		// setup routes
		app := routes.SetupRouter()
		go log.Fatal(app.Listen(":8080"))

	}
}
