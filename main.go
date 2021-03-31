package main

import (
	"log"
	"todo-backend/webapi"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Print("Starting the grocery app")
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*", "http://localhost:3000"}

	r.Use(cors.New(config))
	webapi.Route(r)
	r.Run()
}
