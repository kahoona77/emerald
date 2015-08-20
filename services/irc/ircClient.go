package irc

import (
	"log"

	"github.com/efarrer/iothrottler"
	"github.com/kahoona77/emerald/models"
	"github.com/kahoona77/emerald/services/dataService"
	"github.com/kahoona77/emerald/services/irc/bot"
	"github.com/kahoona77/emerald/services/showsService"
)

// Client is a irc client that cann connect to multiple irc servers
type Client struct {
	DataService  *dataService.DataService   `inject:""`
	ShowsService *showsService.ShowsService `inject:""`

	bots       map[string]*bot.IrcBot
	downloads  map[string]*models.Download
	updateChan chan models.DccUpdate
	connPool   *iothrottler.IOThrottlerPool
}

// NewClient creates a new client
func NewClient() *Client {
	client := new(Client)
	client.bots = make(map[string]*bot.IrcBot)
	client.downloads = make(map[string]*models.Download)
	client.updateChan = make(chan models.DccUpdate)
	client.connPool = iothrottler.NewIOThrottlerPool(iothrottler.Unlimited)

	//start download updates
	go client.updateDownloads()

	return client
}

//ToggleConnection connects or discconnects the client
func (ic *Client) ToggleConnection(server *models.Server) {
	bot := ic.getAndUpdateBot(server)
	if bot.IsConnected() {
		bot.Disconnect()
	} else {
		bot.Connect()
	}
}

//IsServerConnected returns true if the given server is connected
func (ic *Client) IsServerConnected(server *models.Server) bool {
	bot := ic.getAndUpdateBot(server)
	return bot.IsConnected()
}

//GetServerConsole returns the log of the server
func (ic *Client) GetServerConsole(server *models.Server) string {
	bot := ic.getAndUpdateBot(server)
	return bot.GetLog()
}

// GetBot return the Bot of the given serverName
func (ic *Client) GetBot(serverName string) *bot.IrcBot {
	return ic.bots[serverName]
}

func (ic *Client) getAndUpdateBot(server *models.Server) *bot.IrcBot {
	currentBot := ic.bots[server.Name]
	if currentBot == nil {
		// create new bot
		currentBot = bot.NewIrcBot(server, ic.updateChan, ic.connPool, ic.DataService)
		ic.bots[server.Name] = currentBot
	} else {
		//update bot
		currentBot.SetServer(server)
	}
	return currentBot
}

// SetDownloadLimit - Sets the downloadlimit in KiloByte / Second
func (ic *Client) SetDownloadLimit(maxDownStream int) {
	if maxDownStream <= 0 {
		ic.connPool.SetBandwidth(iothrottler.Unlimited)
		log.Printf("download unlimited")
	} else {
		ic.connPool.SetBandwidth(iothrottler.Kbps * iothrottler.Bandwidth(maxDownStream*8))
		log.Printf("currentDownloadLimit: %v", iothrottler.Kbps*iothrottler.Bandwidth(maxDownStream*8))
	}
}
