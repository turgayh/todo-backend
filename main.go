package main

import (
	"log"
	"todo-backend/webapi"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Print("Starting the grocery app")
	r := gin.Default()
	webapi.Route(r)
	r.Run()
}
