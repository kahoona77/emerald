package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kahoona77/emerald/models"
	"github.com/kahoona77/emerald/services/irc"
)

//IrcController -
type IrcController struct {
	IrcClient *irc.Client `inject:""`
}

// ConfigureRoutes configures the routes for this controller
func (ic *IrcController) ConfigureRoutes(route *gin.RouterGroup) {
	route.POST("toggleConnection", ic.toggleConnection)
	route.POST("getServerStatus", ic.getServerStatus)
	route.POST("getServerConsole", ic.getServerConsole)
}

//ToggleConnection toggles the connection of the given server
func (ic *IrcController) toggleConnection(c *gin.Context) {
	var server models.Server
	c.BindJSON(&server)
	ic.IrcClient.ToggleConnection(&server)
	OK(c)
}

//GetServerStatus returns if the server is connected to an IRC HOST
func (ic *IrcController) getServerStatus(c *gin.Context) {
	var server models.Server
	c.BindJSON(&server)
	connected := ic.IrcClient.IsServerConnected(&server)
	renderOk(c, gin.H{"connected": connected})
}

//GetServerConsole returns the console-log of the given server
func (ic *IrcController) getServerConsole(c *gin.Context) {
	var server models.Server
	c.BindJSON(&server)
	console := ic.IrcClient.GetServerConsole(&server)
	renderOk(c, console)
}
