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
	port := flag.Int("port", 8080, "port to serve on")
	logFile := flag.String("log", "emerald.log", "log-file")
	flag.Parse()

	// setup log
	f, err := os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Printf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	var app core.EmeraldApp
	mongoService := mongo.NewService()
	ircClient := irc.NewClient()

	var g inject.Graph
	err = g.Provide(
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
	log.Printf("Emerald started port %d\n", *port)
	fmt.Printf("Emerald started port %d\n", *port)
	addr := fmt.Sprintf(":%d", *port)
	router.Run(addr)
}
