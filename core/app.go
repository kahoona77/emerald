package core

import (
	"github.com/gin-gonic/gin"
	"github.com/kahoona77/emerald/controllers"
	"github.com/kahoona77/emerald/services/jobs"
)

// EmeraldApp is the bas App of Emerald
type EmeraldApp struct {
	DataController      *controllers.DataController      `inject:""`
	ShowsController     *controllers.ShowsController     `inject:""`
	DownloadsController *controllers.DownloadsController `inject:""`
	IrcController       *controllers.IrcController       `inject:""`
	UpdateJob           *jobs.UpdateJob                  `inject:""`
}

//AddControllers add all controllers of emerald to gin
func (app *EmeraldApp) AddControllers(router *gin.Engine) {
	app.DataController.ConfigureRoutes(router.Group("/data"))
	app.ShowsController.ConfigureRoutes(router.Group("/shows"))
	app.DownloadsController.ConfigureRoutes(router.Group("/downloads"))
	app.IrcController.ConfigureRoutes(router.Group("/irc"))
}

//StartJobs starts all jobs
func (app *EmeraldApp) StartJobs() {
	app.UpdateJob.StartJob()
}
