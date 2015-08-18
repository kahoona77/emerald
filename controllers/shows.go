package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kahoona77/emerald/models"
	"github.com/kahoona77/emerald/services/showsService"
)

// ShowsController Creates all routes for the ShowsController
func ShowsController(router *gin.RouterGroup) {
	router.GET("/load", load)
	router.POST("/save", save)
	router.POST("/delete", delete)
	router.GET("/search", search)
	router.GET("/loadEpisodes", loadEpisodes)
	router.GET("/recentEpisodes", recentEpisodes)

}

func load(c *gin.Context) {
	shows, err := showsService.FindAllShows()
	if err != nil {
		renderError(c, err)
	}
	renderOk(c, shows)
}

func save(c *gin.Context) {
	var show models.Show
	c.BindJSON(&show)

	_, err := showsService.UpdateShow(&show)
	if err != nil {
		renderError(c, err)
	}

	renderOk(c, show)
}

func delete(c *gin.Context) {
	var show models.Show
	c.BindJSON(&show)

	err := showsService.DeleteShow(&show)
	if err != nil {
		renderError(c, err)
	}

	renderOk(c, show)
}

func search(c *gin.Context) {
	var query = c.Query("query")

	shows, err := showsService.SearchShow(query)
	if err != nil {
		renderError(c, err)
	}

	renderOk(c, shows)
}

func loadEpisodes(c *gin.Context) {
	var showID = c.Query("showId")
	episodes, err := showsService.LoadEpisodes(showID)
	if err != nil {
		renderError(c, err)
	}

	renderOk(c, episodes)
}

func recentEpisodes(c *gin.Context) {
	var duration = c.Query("duration")
	episodes, err := showsService.LoadRecentEpisodes(duration)
	if err != nil {
		renderError(c, err)
	}

	renderOk(c, episodes)
}
