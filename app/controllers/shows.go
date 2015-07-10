package controllers

import (
	"github.com/kahoona77/emerald/app/models"
	"github.com/kahoona77/emerald/app/services/showsService"
	"github.com/revel/revel"
)

type Shows struct {
	*revel.Controller
}

func (c Shows) Load() revel.Result {
	shows, err := showsService.FindAllShows()
	if err != nil {
		return renderError(c.Controller, err)
	}
	return renderOk(c.Controller, shows)
}

func (c Shows) Save() revel.Result {
	var show models.Show
	readJson(c.Request.Request, &show)

	_, err := showsService.UpdateShow(&show)
	if err != nil {
		return renderError(c.Controller, err)
	}

	return renderOk(c.Controller, show)
}

func (c Shows) Delete() revel.Result {
	var show models.Show
	readJson(c.Request.Request, &show)

	err := showsService.DeleteShow(&show)
	if err != nil {
		return renderError(c.Controller, err)
	}

	return renderOk(c.Controller, show)
}

func (c Shows) Search() revel.Result {
	var query = c.Params.Get("query")

	shows, err := showsService.SearchShow(query)
	if err != nil {
		return renderError(c.Controller, err)
	}

	return renderOk(c.Controller, shows)
}

func (c Shows) LoadEpisodes() revel.Result {
	var showId = c.Params.Get("showId")
	episodes, err := showsService.LoadEpisodes(showId)
	if err != nil {
		return renderError(c.Controller, err)
	}

	return renderOk(c.Controller, episodes)
}

func (c Shows) RecentEpisodes() revel.Result {
	var duration = c.Params.Get("duration")
	episodes, err := showsService.LoadRecentEpisodes(duration)
	if err != nil {
		return renderError(c.Controller, err)
	}

	return renderOk(c.Controller, episodes)
}
