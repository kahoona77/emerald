package models

import (
	"time"
)

//AppConfig the Emerald config
type AppConfig struct {
	Port    int
	LogFile string
	Mongodb string
}

//MongoModel the base model
type MongoModel interface {
	SetId(id string)
	GetId() string
}

// Server represents a irc-server
type Server struct {
	Id       string    `json:"id" bson:"_id"`
	Name     string    `json:"name" bson:"name"`
	Port     int       `json:"port" bson:"port"`
	Channels []Channel `json:"channels" bson:"channels"`
}

func (this *Server) SetId(id string) {
	this.Id = id
}

func (this *Server) GetId() string {
	return this.Id
}

type Channel struct {
	Name string `json:"name" bson:"name"`
}

//Packet
type Packet struct {
	Id       string    `json:"id" bson:"_id"`
	PacketId string    `json:"packetId" bson:"packetId"`
	Size     string    `json:"size" bson:"size"`
	Name     string    `json:"name" bson:"name"`
	Bot      string    `json:"bot" bson:"bot"`
	Channel  string    `json:"channel" bson:"channel"`
	Server   string    `json:"server" bson:"server"`
	Date     time.Time `json:"date" bson:"date"`
}

func NewPacket(packetId string, size string, name string, bot string, channel string, server string, date time.Time) *Packet {
	p := new(Packet)
	p.Id = channel + ":" + bot + ":" + packetId
	p.PacketId = packetId
	p.Size = size
	p.Name = name
	p.Bot = bot
	p.Channel = channel
	p.Server = server
	p.Date = date
	return p
}

func (this *Packet) SetId(id string) {
	this.Id = id
}

func (this *Packet) GetId() string {
	return this.Id
}

// Settings
type EmeraldSettings struct {
	Id            string `json:"id" bson:"_id"`
	Nick          string `json:"nick" bson:"nick"`
	TempDir       string `json:"tempDir" bson:"tempDir"`
	DownloadDir   string `json:"downloadDir" bson:"downloadDir"`
	ShowsFolder   string `json:"showsFolder" bson:"showsFolder"`
	MoviesFolder  string `json:"moviesFolder" bson:"moviesFolder"`
	KodiAddress   string `json:"kodiAddress" bson:"kodiAddress"`
	LogFile       string `json:"logFile" bson:"logFile"`
	MaxDownStream int    `json:"maxDownStream" bson:"maxDownStream"`
}

func (this *EmeraldSettings) SetId(id string) {
	this.Id = id
}

func (this *EmeraldSettings) GetId() string {
	return this.Id
}

//Show
type Show struct {
	Id         string    `json:"id" bson:"_id"`
	Name       string    `json:"name" bson:"name"`
	Banner     string    `json:"banner" bson:"banner"`
	Poster     string    `json:"poster" bson:"poster"`
	FirstAired time.Time `json:"firstAired" bson:"firstAired"`
	Overview   string    `json:"overview" bson:"overview"`
	SearchName string    `json:"searchName" bson:"searchName"`
	Folder     string    `json:"folder" bson:"folder"`
}

func (this *Show) SetId(id string) {
	this.Id = id
}

func (this *Show) GetId() string {
	return this.Id
}

//Episode
type Episode struct {
	Id            string    `json:"id" bson:"_id"`
	ShowId        string    `json:"shpwId" bson:"showId"`
	Name          string    `json:"name" bson:"name"`
	FirstAired    time.Time `json:"firstAired" bson:"firstAired"`
	Overview      string    `json:"overview" bson:"overview"`
	Filename      string    `json:"filename" bson:"filename"`
	EpisodeNumber uint64    `json:"episodeNumber" bson:"episodeNumber"`
	SeasonNumber  uint64    `json:"seasonNumber" bson:"seasonNumber"`
}

func (this *Episode) SetId(id string) {
	this.Id = id
}

func (this *Episode) GetId() string {
	return this.Id
}

//RecentEpisode
type RecentEpisode struct {
	Episode Episode `json:"episode"`
	Show    Show    `json:"show"`
}
