package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kahoona77/emerald/models"
	"github.com/kahoona77/emerald/services/filesService"
	"github.com/kahoona77/emerald/services/irc"
)

//DownloadsController -
type DownloadsController struct {
	IrcClient    *irc.Client                `inject:""`
	FilesService *filesService.FilesService `inject:""`
}

//DownloadPacket starts the download of a packet
func (dc *DownloadsController) DownloadPacket(c *gin.Context) {
	var packet models.Packet
	c.BindJSON(&packet)
	dc.IrcClient.DownloadPacket(&packet)
	OK(c)
}

//ListDownloads loads all downloads
func (dc *DownloadsController) ListDownloads(c *gin.Context) {
	downloads := dc.IrcClient.ListDownloads()
	renderOk(c, downloads)
}

//StopDownload stops the given Download
func (dc *DownloadsController) StopDownload(c *gin.Context) {
	var download models.Download
	c.BindJSON(&download)
	dc.IrcClient.StopDownload(&download)
	OK(c)
}

//CancelDownload canceles the given download
func (dc *DownloadsController) CancelDownload(c *gin.Context) {
	var download models.Download
	c.BindJSON(&download)
	dc.IrcClient.CancelDownload(&download)
	OK(c)
}

//ResumeDownload resumes the given download
func (dc *DownloadsController) ResumeDownload(c *gin.Context) {
	var download models.Download
	c.BindJSON(&download)
	dc.IrcClient.ResumeDownload(&download)
	OK(c)
}

//LoadFiles loads all files from the files-service
func (dc *DownloadsController) LoadFiles(c *gin.Context) {
	files := dc.FilesService.GetFiles()
	renderOk(c, files)
}

//DeleteFiles deletes the given files
func (dc *DownloadsController) DeleteFiles(c *gin.Context) {
	var files []filesService.File
	c.BindJSON(&files)
	err := dc.FilesService.DeleteFiles(files)
	if err != nil {
		log.Printf("ERROR: %v", err)
		renderErrorMsg(c, "Error while deleting files: "+err.Error())
		return
	}
	OK(c)
}

//MoveFilesToMovies moves the given Files to the Movied-Dir
func (dc *DownloadsController) MoveFilesToMovies(c *gin.Context) {
	var files []filesService.File
	c.BindJSON(&files)
	err := dc.FilesService.MoveFilesToMovies(files)
	if err != nil {
		log.Printf("ERROR: %v", err)
		renderErrorMsg(c, "Error while moving files: "+err.Error())
		return
	}
	OK(c)
}
