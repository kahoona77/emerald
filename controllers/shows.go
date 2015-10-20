package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kahoona77/emerald/models"
	"github.com/kahoona77/emerald/services/dataService"
	"github.com/kahoona77/emerald/services/showsService"
)

// ShowsController Creates all routes for the ShowsController
type ShowsController struct {
	ShowsService *showsService.ShowsService `inject:""`
	DataService  *dataService.DataService   `inject:""`
}

// ConfigureRoutes configures the routes for this controller
func (sc *ShowsController) ConfigureRoutes(route *gin.RouterGroup) {
	route.GET("load", sc.load)
	route.POST("save", sc.save)
	route.GET("delete", sc.delete)
	route.GET("search", sc.search)
	route.GET("loadEpisodes", sc.loadEpisodes)
	route.GET("recentEpisodes", sc.recentEpisodes)
	route.GET("updateEpisodes", sc.updateEpisodes)
}

//Load loads all shows
func (sc *ShowsController) load(c *gin.Context) {
	shows, err := sc.ShowsService.FindAllShows()
	if err != nil {
		renderError(c, err)
	}
	renderOk(c, shows)
}

//Save saves the given show
func (sc *ShowsController) save(c *gin.Context) {
	var show models.Show
	c.BindJSON(&show)

	_, err := sc.ShowsService.UpdateShow(&show)
	if err != nil {
		renderError(c, err)
	}

	renderOk(c, show)
}

//Delete deletes the given show
func (sc *ShowsController) delete(c *gin.Context) {
	var show models.Show
	c.BindJSON(&show)

	err := sc.ShowsService.DeleteShow(&show)
	if err != nil {
		renderError(c, err)
	}

	renderOk(c, show)
}

//Search searches for shows
func (sc *ShowsController) search(c *gin.Context) {
	var query = c.Query("query")

	shows, err := sc.ShowsService.SearchShow(query)
	if err != nil {
		renderError(c, err)
	}

	renderOk(c, shows)
}

//LoadEpisodes loads all episodes of a show
func (sc *ShowsController) loadEpisodes(c *gin.Context) {
	var showID = c.Query("showId")
	episodes, err := sc.ShowsService.LoadEpisodes(showID)
	if err != nil {
		renderError(c, err)
	}

	renderOk(c, episodes)
}

//RecentEpisodes loads all recent epiode of a show
func (sc *ShowsController) recentEpisodes(c *gin.Context) {
	var duration = c.Query("duration")
	episodes, err := sc.ShowsService.LoadRecentEpisodes(duration)
	if err != nil {
		renderError(c, err)
	}

	renderOk(c, episodes)
}

//UpdateEpisodes update
func (sc *ShowsController) updateEpisodes(c *gin.Context) {
	settings := sc.DataService.LoadSettings()
	sc.ShowsService.ScanDownloadDir(settings)
	OK(c)
}
