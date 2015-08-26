package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/facebookgo/inject"
	"github.com/gin-gonic/gin"
	"github.com/kahoona77/emerald/core"
	"github.com/kahoona77/emerald/services/irc"
	"github.com/kahoona77/emerald/services/mongo"
)

func main() {
	// command line flags
	confFile := flag.String("conf", "emerald.conf", "location of config-file")
	flag.Parse()
	conf := core.LoadConfiguration(*confFile)

	var app core.EmeraldApp
	mongoService := mongo.NewService(&conf)
	ircClient := irc.NewClient()

	var g inject.Graph
	err := g.Provide(
		&inject.Object{Value: &conf},
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
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	//Middlewares
	// router.Use(gin.Logger())
	router.Use(gin.Recovery())

	//static files
	router.Static("/assets", "./assets")
	router.StaticFile("/", "./assets/index.html")

	//register all controllers
	app.AddControllers(router)

	//Start all jobs
	app.StartJobs()

	// Listen and server on 0.0.0.0:8080
	addr := fmt.Sprintf(":%d", conf.Port)
	log.Printf("Emerald started port %v\n", addr)
	fmt.Printf("Emerald started port %v\n", addr)
	router.Run(addr)
}
