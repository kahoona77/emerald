package main

import (
	"fmt"
	"os"

	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/kahoona77/emerald/core"
	"github.com/kahoona77/emerald/services/irc"
	"github.com/kahoona77/emerald/services/mongo"
)

func main() {
	var app core.EmeraldApp
	mongoService := mongo.NewService()
	ircClient := irc.NewClient()

	var g inject.Graph
	err := g.Provide(
		&inject.Object{Value: &app},
		&inject.Object{Value: mongoService},
		&inject.Object{Value: ircClient},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := g.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	//setup gin
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFile("/", "./assets/index.html")

	//register all controllers
	app.AddControllers(router)

	// Listen and server on 0.0.0.0:8080
	router.Run(":8080")
}
