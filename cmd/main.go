package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/matilloret/cmd/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		log.Fatal("PORT is not set in .env file")
	}

	app := api.NewAPIServer(port)

	if err := app.Run(); err != nil {
		log.Fatal("error to inicialize the server", err.Error())
	}
}
