package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	api "app/api"
)

func main() {

	// Get env
	env, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	api.Start(env)
}