package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kahoona77/emerald/models"
	"github.com/kahoona77/emerald/services/dataService"
	"github.com/revel/revel"
)

// DataController creates all routes for the DataController
func DataController(router *gin.RouterGroup) {
	router.GET("/loadSettings", loadSettings)
	router.POST("/saveSettings", saveSettings)
	router.GET("/countPackets", countPackets)
	router.GET("/findPackets", findPackets)
	router.GET("/loadServers", loadServers)
	router.POST("/deleteServer", deleteServer)
	router.POST("/saveServer", saveServer)
}

func saveServer(c *gin.Context) {
	var server models.Server
	c.BindJSON(&server)

	err := dataService.SaveServer(&server)
	if err != nil {
		renderErrorMsg(c, "Error while saving server")
	}

	renderOk(c, server)
}

func deleteServer(c *gin.Context) {
	var server models.Server
	c.BindJSON(&server)

	err := dataService.DeleteServer(&server)
	if err != nil {
		renderErrorMsg(c, "Error while deleting server")
	}

	renderOk(c, server)
}

func loadServers(c *gin.Context) {
	servers, _ := dataService.FindAllServers()
	renderOk(c, servers)
}

func loadSettings(c *gin.Context) {
	settings := dataService.LoadSettings()
	renderOk(c, settings)
}

func saveSettings(c *gin.Context) {
	var settings models.XtvSettings
	c.BindJSON(&settings)

	dataService.SaveSettings(&settings)
	renderOk(c, nil)
}

func findPackets(c *gin.Context) {
	var query = c.Query("query")
	packtes, err := dataService.FindPackets(query)
	if err != nil {
		renderError(c, err)
	}
	renderOk(c, packtes)
}

func countPackets(c *gin.Context) {
	deletedPacktes, _ := dataService.DeleteOldPackets()
	revel.INFO.Printf("Deleted %v packets", deletedPacktes)

	packetCount, _ := dataService.CountPackets()
	renderOk(c, packetCount)
}
