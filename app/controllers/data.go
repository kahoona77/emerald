package controllers

import (
	"github.com/kahoona77/emerald/app/models"
	"github.com/kahoona77/emerald/app/services/dataService"
	"github.com/revel/revel"
)

type Data struct {
	*revel.Controller
}

func (c Data) SaveServer() revel.Result {
	var server models.Server
	readJson(c.Request.Request, &server)

	err := dataService.SaveServer(&server)
	if err != nil {
		return c.RenderJson(ERROR("Error while saving server"))
	}

	return renderOk(c.Controller, server)
}

func (c Data) DeleteServer() revel.Result {
	var server models.Server
	readJson(c.Request.Request, &server)

	err := dataService.DeleteServer(&server)
	if err != nil {
		return c.RenderJson(ERROR("Error while deleting server"))
	}

	return c.RenderJson(OK())
}

func (c Data) LoadServers() revel.Result {
	servers, _ := dataService.FindAllServers()
	return renderOk(c.Controller, servers)
}

func (c Data) LoadSettings() revel.Result {
	settings := dataService.LoadSettings()
	return renderOk(c.Controller, settings)
}

func (c Data) SaveSettings() revel.Result {
	var settings models.XtvSettings
	readJson(c.Request.Request, &settings)

	dataService.SaveSettings(&settings)
	return c.RenderJson(OK())
}

func (c Data) FindPackets() revel.Result {
	var query = c.Params.Get("query")
	packtes, err := dataService.FindPackets(query)
	if err != nil {
		return renderError(c.Controller, err)
	}
	return renderOk(c.Controller, packtes)
}

func (c Data) CountPackets() revel.Result {
	deletedPacktes, _ := dataService.DeleteOldPackets()
	revel.INFO.Printf("Deleted %v packets", deletedPacktes)

	packetCount, _ := dataService.CountPackets()
	return renderOk(c.Controller, packetCount)
}
