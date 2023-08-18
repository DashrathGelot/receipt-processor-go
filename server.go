package main

import (
	"log"
	"receipt-processor/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("=======================================================")
	log.Println("=====  Fetch Backend Engineer Assessment 2023  ========")
	log.Println("=======================================================")

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	routes.ReceiptRoute(router)

	log.Println("API is running on localhost:8080")
	router.Run()
}
