package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kahoona77/emerald/controllers"
	"github.com/kahoona77/emerald/services/mongo"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFile("/", "./assets/index.html")

	controllers.DataController(router.Group("/data"))
	controllers.ShowsController(router.Group("/shows"))

	mongo.InitDB()

	// Listen and server on 0.0.0.0:8080
	router.Run(":8080")
}
