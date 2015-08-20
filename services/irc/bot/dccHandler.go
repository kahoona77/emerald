package bot

import (
	"log"
	"net"
	"strconv"
	"strings"

	irc "github.com/fluffle/goirc/client"
	"github.com/kahoona77/emerald/models"
)

//HandleDCC handles all incomming DCC requests
func (ib *IrcBot) handleDCC(conn *irc.Conn, line *irc.Line) {
	request := strings.Split(line.Args[2], " ")
	ctcpType := line.Args[0]
	settings := ib.dataService.LoadSettings()
	if ctcpType == "DCC" {
		cmd := request[0]
		if cmd == "SEND" {
			ib.handleSend(request, conn, line, settings)
		} else if cmd == "ACCEPT" {
			ib.handleAccept(request, settings)
		} else {
			log.Printf("received unmatched DCC command: %v", cmd)
		}
	}
}

func (ib *IrcBot) handleSend(request []string, conn *irc.Conn, line *irc.Line, settings *models.XtvSettings) {
	fileName := request[1]
	addrInt, _ := strconv.ParseInt(request[2], 0, 64)
	address := inetNtoa(addrInt)
	port := request[3]
	size, _ := strconv.ParseInt(request[4], 0, 64)

	log.Printf("received SEND - file: %v, addr: %v, port: %v, size:%v\n", fileName, address.String(), port, size)
	fileEvent := models.DccFileEvent{"SEND", fileName, address, port, size}

	resume, startPos := fileExists(&fileEvent, settings)

	if resume {
		// file already exists -> send resume request
		msg := fileName + " " + port + " " + strconv.FormatInt(startPos, 10)
		log.Printf("sending resume [%v]", msg)
		conn.Ctcp(line.Nick, "DCC RESUME", msg)
		//add to resumes
		ib.resumes[fileEvent.FileName] = &fileEvent
	} else {
		// This is a new file start from beginning
		go ib.startDownload(&fileEvent, startPos, settings)
	}
}

func (ib *IrcBot) handleAccept(request []string, settings *models.XtvSettings) {
	log.Printf("received ACCEPT")

	fileName := request[1]
	//port := request[2]
	position, err := strconv.ParseInt(request[3], 10, 64)

	if err != nil {
		log.Printf("error while parsing position %v", err)
		return
	}

	//find resume
	fileEvent := ib.resumes[fileName]
	delete(ib.resumes, fileName)
	if fileEvent == nil {
		log.Printf("can not find resume for %v", fileName)
		return
	}

	//start the download in new goroutine
	go ib.startDownload(fileEvent, position, settings)
}

// Convert uint to net.IP
func inetNtoa(ipnr int64) net.IP {
	var bytes [4]byte
	bytes[0] = byte(ipnr & 0xFF)
	bytes[1] = byte((ipnr >> 8) & 0xFF)
	bytes[2] = byte((ipnr >> 16) & 0xFF)
	bytes[3] = byte((ipnr >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}
