package controllers

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/kahoona77/emerald/models"
	"github.com/kahoona77/emerald/services/dataService"
	"github.com/kahoona77/emerald/services/irc"
)

// DataController creates all routes for the DataController
type DataController struct {
	DataService *dataService.DataService `inject:""`
	IrcClient   *irc.Client              `inject:""`
	Config      *models.AppConfig        `inject:""`
}

// SaveServer saves a server
func (dc *DataController) SaveServer(c *gin.Context) {
	var server models.Server
	c.BindJSON(&server)

	err := dc.DataService.SaveServer(&server)
	if err != nil {
		renderErrorMsg(c, "Error while saving server")
	}

	renderOk(c, server)
}

//DeleteServer deletes a server
func (dc *DataController) DeleteServer(c *gin.Context) {
	var server models.Server
	c.BindJSON(&server)

	err := dc.DataService.DeleteServer(&server)
	if err != nil {
		renderErrorMsg(c, "Error while deleting server")
	}

	renderOk(c, server)
}

//LoadServers loads all servers
func (dc *DataController) LoadServers(c *gin.Context) {
	servers, _ := dc.DataService.FindAllServers()
	renderOk(c, servers)
}

//LoadSettings loads the settings
func (dc *DataController) LoadSettings(c *gin.Context) {
	settings := dc.DataService.LoadSettings()
	renderOk(c, settings)
}

//SaveSettings saves the settings
func (dc *DataController) SaveSettings(c *gin.Context) {
	var settings models.XtvSettings
	c.BindJSON(&settings)

	dc.DataService.SaveSettings(&settings)
	dc.IrcClient.SetDownloadLimit(settings.MaxDownStream)
	renderOk(c, nil)
}

//FindPackets finds DCC packets
func (dc *DataController) FindPackets(c *gin.Context) {
	var query = c.Query("query")
	packtes, err := dc.DataService.FindPackets(query)
	if err != nil {
		renderError(c, err)
	}
	renderOk(c, packtes)
}

//CountPackets counts alls packest in the DB - first delets old packets
func (dc *DataController) CountPackets(c *gin.Context) {
	deletedPacktes, _ := dc.DataService.DeleteOldPackets()
	log.Printf("Deleted %v packets", deletedPacktes)

	packetCount, _ := dc.DataService.CountPackets()
	renderOk(c, packetCount)
}

//LoadLogFile loads the logFile from Disk
func (dc *DataController) LoadLogFile(c *gin.Context) {
	logFile := dc.Config.LogFile
	buf, _ := ioutil.ReadFile(logFile)
	renderOk(c, string(buf))
}

//ClearLogFile clears the logFile
func (dc *DataController) ClearLogFile(c *gin.Context) {
	logFile := dc.Config.LogFile
	ioutil.WriteFile(logFile, []byte(""), 0644)
	OK(c)
}
