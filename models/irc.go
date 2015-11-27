package models

import (
	"net"
	"time"
)

// DccFileEvent -
type DccFileEvent struct {
	Type     string
	FileName string
	IP       net.IP
	Port     string
	Size     int64
}

// UpdateType -
type UpdateType int

// Update constants
const (
	UPDATE   UpdateType = 0
	COMPLETE UpdateType = 1
	FAIL     UpdateType = 2
)

// DccUpdate -
type DccUpdate struct {
	File       string
	TotalBytes int64
	Size       int64
	Type       UpdateType
}

//Download represent a Download in Emerald
type Download struct {
	ID            string    `json:"id"`
	Status        string    `json:"status"`
	File          string    `json:"file"`
	PacketID      string    `json:"packetId"`
	Server        string    `json:"server"`
	Bot           string    `json:"bot"`
	BytesReceived int64     `json:"bytesReceived"`
	Size          int64     `json:"size"`
	Speed         float32   `json:"speed"`
	Remaining     int64     `json:"remaining"`
	LastUpdate    time.Time `json:"-"`
}

//Download represent a Download in Emerald
type DirectDownload struct {
	Server  string `json:"server"`
	Message string `json:"message"`
}
