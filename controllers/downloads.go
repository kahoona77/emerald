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

// ConfigureRoutes configures the routes for this controller
func (dc *DownloadsController) ConfigureRoutes(route *gin.RouterGroup) {
	route.POST("downloadPacket", dc.downloadPacket)
	route.GET("listDownloads", dc.listDownloads)
	route.POST("stopDownload", dc.stopDownload)
	route.POST("cancelDownload", dc.cancelDownload)
	route.POST("resumeDownload", dc.resumeDownload)
	route.GET("loadFiles", dc.loadFiles)
	route.POST("deleteFiles", dc.deleteFiles)
	route.POST("moveFilesToMovies", dc.moveFilesToMovies)
}

//DownloadPacket starts the download of a packet
func (dc *DownloadsController) downloadPacket(c *gin.Context) {
	var packet models.Packet
	c.BindJSON(&packet)
	dc.IrcClient.DownloadPacket(&packet)
	OK(c)
}

//ListDownloads loads all downloads
func (dc *DownloadsController) listDownloads(c *gin.Context) {
	downloads := dc.IrcClient.ListDownloads()
	renderOk(c, downloads)
}

//StopDownload stops the given Download
func (dc *DownloadsController) stopDownload(c *gin.Context) {
	var download models.Download
	c.BindJSON(&download)
	dc.IrcClient.StopDownload(&download)
	OK(c)
}

//CancelDownload canceles the given download
func (dc *DownloadsController) cancelDownload(c *gin.Context) {
	var download models.Download
	c.BindJSON(&download)
	dc.IrcClient.CancelDownload(&download)
	OK(c)
}

//ResumeDownload resumes the given download
func (dc *DownloadsController) resumeDownload(c *gin.Context) {
	var download models.Download
	c.BindJSON(&download)
	dc.IrcClient.ResumeDownload(&download)
	OK(c)
}

//LoadFiles loads all files from the files-service
func (dc *DownloadsController) loadFiles(c *gin.Context) {
	files := dc.FilesService.GetFiles()
	renderOk(c, files)
}

//DeleteFiles deletes the given files
func (dc *DownloadsController) deleteFiles(c *gin.Context) {
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
func (dc *DownloadsController) moveFilesToMovies(c *gin.Context) {
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
