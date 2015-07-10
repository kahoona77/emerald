package dataService

import (
	"github.com/kahoona77/emerald/app/models"
	"github.com/kahoona77/emerald/app/services/mongo"
)

func FindAllServers() ([]models.Server, error) {
	var results []models.Server
	err := mongo.All("servers", &results)
	return results, err
}

func DeleteServer(server *models.Server) error {
	err := mongo.Remove("servers", server.Id)
	return err
}

func SaveServer(server *models.Server) error {
	_, err := mongo.Save("servers", server.Id, server)
	return err
}

func LoadSettings() *models.XtvSettings {
	var settings models.XtvSettings
	mongo.FindFirst("settings", &settings)
	return &settings
}

func SaveSettings(settings *models.XtvSettings) error {
	_, err := mongo.Save("settings", settings.Id, settings)
	return err
}
