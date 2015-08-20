package irc

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/kahoona77/emerald/models"
)

//DownloadFromPacket creates a new Download from a Packet
func DownloadFromPacket(packet *models.Packet) *models.Download {
	d := models.Download{ID: packet.Name, Status: "WAITING", File: packet.Name, PacketID: packet.PacketId, Bot: packet.Bot, Server: packet.Server}
	d.LastUpdate = time.Now()
	return &d
}

//DownloadPacket starts teh Download of a Packet
func (ic *Client) DownloadPacket(packet *models.Packet) {
	bot := ic.GetBot(packet.Server)
	download := DownloadFromPacket(packet)
	ic.downloads[download.ID] = download
	bot.StartDownload(download)
}

//ListDownloads returs an array of all Downloads
func (ic *Client) ListDownloads() []*models.Download {
	v := make([]*models.Download, 0, len(ic.downloads))

	for _, value := range ic.downloads {
		v = append(v, value)
	}
	return v
}

//StopDownload stops the given download
func (ic *Client) StopDownload(download *models.Download) {
	bot := ic.GetBot(download.Server)
	bot.StopDownload(download)
}

//CancelDownload cancels the given download
func (ic *Client) CancelDownload(parsedDownload *models.Download) {
	download := ic.downloads[parsedDownload.ID]
	if download != nil {
		if download.Status == "RUNNING" {
			ic.StopDownload(download)
		}
		delete(ic.downloads, download.ID)
	}
}

//ResumeDownload resumes the given download
func (ic *Client) ResumeDownload(parsedDownload *models.Download) {
	download := ic.downloads[parsedDownload.ID]
	if download != nil {
		if download.Status != "RUNNING" {
			bot := ic.GetBot(download.Server)
			bot.StartDownload(download)
		}
	}
}

func (ic *Client) updateDownloads() {
	for {
		update := <-ic.updateChan
		switch update.Type {
		case models.UPDATE:
			ic.updateDownload(update)
		case models.COMPLETE:
			ic.completeDownload(update.File)
		case models.FAIL:
			ic.failDownload(update.File)
		}
	}
}

func (ic *Client) updateDownload(update models.DccUpdate) {
	download := ic.downloads[update.File]
	if download != nil {
		//calc speed
		now := time.Now()
		sizeDelta := (update.TotalBytes - download.BytesReceived) / 1024
		timeDelta := (now.UnixNano() - download.LastUpdate.UnixNano())
		download.Speed = (float32(sizeDelta) / float32(timeDelta)) * 1000 * 1000 * 1000

		//update download
		download.LastUpdate = now
		download.Status = "RUNNING"
		download.BytesReceived = update.TotalBytes
		download.Size = update.Size
	} else {
		log.Printf("download not found: %v in %v", update.File, ic.downloads)
	}
}

func (ic *Client) completeDownload(file string) {
	download := ic.downloads[file]
	if download != nil {
		log.Printf("Download completed '%v'", download.File)
		download.Status = "COMPLETE"

		//move file to destination
		settings := ic.DataService.LoadSettings()
		srcFile := filepath.FromSlash(settings.TempDir + "/" + file)
		absoluteFile := settings.DownloadDir + "/" + file
		destFile := filepath.FromSlash(absoluteFile)
		err := os.Rename(srcFile, destFile)
		if err != nil {
			log.Printf("Error while moving file to destination: %s", err)
		}

		//start smart episode matching
		ic.ShowsService.MoveEpisode(absoluteFile, settings, true)

	} else {
		log.Printf("download not found: %v in %v", file, ic.downloads)
	}
}

func (ic *Client) failDownload(file string) {
	download := ic.downloads[file]
	if download != nil {
		log.Printf("Download failed '%v'", download.File)
		download.Status = "FAILED"

		//TODO what todo with file

	} else {
		log.Printf("download not found: %v in %v", file, ic.downloads)
	}
}
