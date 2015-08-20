package dataService

import (
	"strings"
	"time"

	"log"

	"github.com/kahoona77/emerald/models"
	"github.com/kahoona77/emerald/services/mongo"
	"labix.org/v2/mgo/bson"
)

const serversCollection = "servers"
const settingsCollection = "settings"
const packetsCollection = "packets"

// DataService -
type DataService struct {
	MongoService *mongo.MongoService `inject:""`
}

func (d *DataService) FindAllServers() ([]models.Server, error) {
	var results []models.Server
	err := d.MongoService.All(serversCollection, &results)
	return results, err
}

func (d *DataService) DeleteServer(server *models.Server) error {
	err := d.MongoService.Remove(serversCollection, server.Id)
	return err
}

func (d *DataService) SaveServer(server *models.Server) error {
	_, err := d.MongoService.Save(serversCollection, server.Id, server)
	return err
}

func (d *DataService) LoadSettings() *models.XtvSettings {
	var settings models.XtvSettings
	d.MongoService.FindFirst(settingsCollection, &settings)
	return &settings
}

func (d *DataService) SaveSettings(settings *models.XtvSettings) error {
	_, err := d.MongoService.Save(settingsCollection, settings.Id, settings)
	return err
}

func (d *DataService) DeleteOldPackets() (int, error) {
	minusOneDay, _ := time.ParseDuration("-24h")
	yesterday := time.Now().Add(minusOneDay)
	removeQuery := bson.M{"date": bson.M{"$lt": yesterday}}

	info, err := d.MongoService.RemoveAll(packetsCollection, &removeQuery)
	log.Printf("Deleted:  %v Packets \n", info.Removed)
	return info.Removed, err
}

func (d *DataService) CountPackets() (int, error) {
	return d.MongoService.CountAll(packetsCollection)
}

// FindPackets finds all packets matching the query
func (d *DataService) FindPackets(query string) ([]models.Packet, error) {
	queryRegex := d.createRegexQuery(query)
	queryObject := bson.M{"name": bson.M{"$regex": queryRegex, "$options": "i"}}

	var packets []models.Packet
	err := d.MongoService.FindWithQuery(packetsCollection, &queryObject, &packets)
	return packets, err
}

func (d *DataService) createRegexQuery(query string) string {
	parts := strings.Split(query, " ")
	return strings.Join(parts, ".*")
}

func (d *DataService) SavePacket(packet *models.Packet) error {
	_, err := d.MongoService.Save(packetsCollection, packet.Id, packet)
	return err
}
