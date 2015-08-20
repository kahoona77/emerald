package bot

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/efarrer/iothrottler"
	irc "github.com/fluffle/goirc/client"
	"github.com/kahoona77/emerald/models"
	"github.com/kahoona77/emerald/services/dataService"
)

// IrcBot can connect to a specidic server
type IrcBot struct {
	server      *models.Server
	updateChan  chan models.DccUpdate
	connPool    *iothrottler.IOThrottlerPool
	dataService *dataService.DataService

	conn       *irc.Conn
	consoleLog []string
	regex      *regexp.Regexp
	logCount   int
	resumes    map[string]*models.DccFileEvent
}

//NewIrcBot creates a new bot
func NewIrcBot(server *models.Server, updateChan chan models.DccUpdate, pool *iothrottler.IOThrottlerPool, dataService *dataService.DataService) *IrcBot {
	bot := new(IrcBot)
	bot.server = server
	bot.updateChan = updateChan
	bot.connPool = pool
	bot.dataService = dataService

	//Initialize internal stuff
	bot.consoleLog = make([]string, 0)
	bot.regex, _ = regexp.Compile(`(#[0-9]+).*\[\s*([0-9|\.]+[BbGgiKMs]+)\]\s+(.+).*`)
	bot.resumes = make(map[string]*models.DccFileEvent)
	return bot
}

func (ib *IrcBot) SetServer(server *models.Server) {
	ib.server = server
}

//IsConnected checks whether a bot is connected to its server
func (ib *IrcBot) IsConnected() bool {
	if ib.conn == nil {
		return false
	}
	return ib.conn.Connected()
}

// Connect connects the bot to its serve
func (ib *IrcBot) Connect() {
	settings := ib.dataService.LoadSettings()
	// create a config and fiddle with it first:
	cfg := irc.NewConfig(settings.Nick)
	cfg.Server = ib.server.Name + ":" + strconv.Itoa(ib.server.Port)
	ib.conn = irc.Client(cfg)

	// Join channels
	ib.conn.HandleFunc("connected",
		func(conn *irc.Conn, line *irc.Line) {
			ib.logToConsole("connected to " + ib.server.Name + ":" + strconv.Itoa(ib.server.Port))

			for _, channel := range ib.server.Channels {
				ib.logToConsole("joining channel " + channel.Name)
				conn.Join(channel.Name)
			}
		})

	// Parse Messages
	ib.conn.HandleFunc("PRIVMSG", ib.parseMessage)

	ib.conn.HandleFunc("372", ib.log372)

	ib.conn.HandleFunc("DISCONNECTED", ib.reconnect)

	ib.conn.HandleFunc("CTCP", ib.handleDCC)

	// Tell client to connect.
	log.Printf("Connecting to '%v'", ib.server.Name)
	if err := ib.conn.Connect(); err != nil {
		log.Printf("Connection error: %v\n", err)
		ib.logToConsole("Connection error: " + err.Error())
	}
}

// Disconnect disconnects the bot from its server. Currently NOOP
func (ib *IrcBot) Disconnect() {
	// TODO
	//ib.conn.shutdown()
}

func (ib *IrcBot) reconnect(conn *irc.Conn, line *irc.Line) {
	log.Printf("Discconected from '%v'. Reconnecting now ...", ib.server.Name)
	ib.Connect()
}

func (ib *IrcBot) log372(conn *irc.Conn, line *irc.Line) {
	ib.logToConsole(line.Text())
}

func (ib *IrcBot) parseMessage(conn *irc.Conn, line *irc.Line) {
	ib.parsePacket(conn, line)
}

func (ib *IrcBot) parsePacket(conn *irc.Conn, line *irc.Line) *models.Packet {
	result := ib.regex.FindStringSubmatch(line.Text())
	if result == nil {
		return nil
	}

	fileName := cleanFileName(result[3])
	packet := models.NewPacket(result[1], result[2], fileName, line.Nick, line.Target(), ib.server.Name, line.Time)

	//save packet
	if packet != nil {
		ib.dataService.SavePacket(packet)
	}

	return packet
}

func cleanFileName(filename string) string {
	return strings.Trim(filename, "\u263B\u263C\u0002\u000f ")
}

func (ib *IrcBot) logToConsole(msg string) {
	if ib.logCount > 500 {
		ib.consoleLog = make([]string, 0)
		ib.logCount = 0
	}
	ib.consoleLog = append(ib.consoleLog, msg)
	ib.logCount++
}

//GetLog returns the consoleLog of this bot
func (ib *IrcBot) GetLog() string {
	return strings.Join(ib.consoleLog, "\n")
}

// StartDownload starts the given download
func (ib *IrcBot) StartDownload(download *models.Download) {
	ib.logToConsole("Starting Download: " + download.File)
	msg := "xdcc send " + getCleanPacketID(download)
	ib.conn.Privmsg(download.Bot, msg)
}

//StopDownload stops the given download
func (ib *IrcBot) StopDownload(download *models.Download) {
	ib.logToConsole("Stopping Download: " + download.File)
	msg := "xdcc cancel"
	ib.conn.Privmsg(download.Bot, msg)
}

func getCleanPacketID(download *models.Download) string {
	return strings.Replace(download.PacketID, "#", "", -1)
}
