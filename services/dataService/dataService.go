package dataService

import (
	"strings"
	"time"

	"github.com/kahoona77/emerald/models"
	"github.com/kahoona77/emerald/services/mongo"
	"labix.org/v2/mgo/bson"
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

func DeleteOldPackets() (int, error) {
	minusOneDay, _ := time.ParseDuration("-24h")
	yesterday := time.Now().Add(minusOneDay)
	removeQuery := bson.M{"date": bson.M{"$lt": yesterday}}

	info, err := mongo.RemoveAll("packets", &removeQuery)
	return info.Removed, err
}

func CountPackets() (int, error) {
	return mongo.CountAll("packets")
}

// FindPackets finds all packets matching the query
func FindPackets(query string) ([]models.Packet, error) {
	queryRegex := createRegexQuery(query)
	queryObject := bson.M{"name": bson.M{"$regex": queryRegex, "$options": "i"}}

	var packets []models.Packet
	err := mongo.FindWithQuery("packets", &queryObject, &packets)
	return packets, err
}

func createRegexQuery(query string) string {
	parts := strings.Split(query, " ")
	return strings.Join(parts, ".*")
}
