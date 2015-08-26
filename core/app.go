package core

import (
	"reflect"
	"unicode"
	"unicode/utf8"

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
	app.addController(router.Group("/data"), app.DataController)
	app.addController(router.Group("/shows"), app.ShowsController)
	app.addController(router.Group("/downloads"), app.DownloadsController)
	app.addController(router.Group("/irc"), app.IrcController)
}

func (app *EmeraldApp) addController(route *gin.RouterGroup, controller interface{}) {
	c := reflect.ValueOf(controller)
	typeOfT := c.Type()
	for i := 0; i < c.NumMethod(); i++ {
		actionName := lowerFirst(typeOfT.Method(i).Name)
		actionInterface := c.Method(i).Interface()
		switch actionMethod := actionInterface.(type) {
		case (func(c *gin.Context)):
			route.Any(actionName, actionMethod)
		}
	}
}

func lowerFirst(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

//StartJobs starts all jobs
func (app *EmeraldApp) StartJobs() {
	app.UpdateJob.StartJob()
}
