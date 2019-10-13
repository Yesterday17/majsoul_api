package majsoul_api

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	ApiVersion          = "0.6.58.w"
	LatestTestedVersion = "0.6.73.w"

	MajsoulCN = "www.majsoul.com/1"
	MajsoulJP = "game.mahjongsoul.com"
	MajsoulEN = "mahjongsoul.game.yo-star.com"
)

type MajsoulAPI struct {
	Base   string
	Region string
	Secret string

	ignoreVersionCheck     bool
	panicOnVersionMismatch bool

	lobbyClient *lobbyClient
	gameClient  *gameClient

	systemEmailUrl     string
	resourceVersion    *map[string]prefixVersion
	regions            map[string]string
	activeWebSocketUrl string
	websocketServers   []string
}

// NewMajsoulAPI Create a new majsoul api instance
func NewMajsoulAPI(base string) (api *MajsoulAPI, err error) {
	api = &MajsoulAPI{
		Base:                   base,
		Region:                 "mainland",
		Secret:                 "lailai",
		ignoreVersionCheck:     false,
		panicOnVersionMismatch: true,
	}

	// init config.js
	err = api.initConfig()
	if err != nil {
		return
	}

	// init recommend_list
	err = api.initWebSocket()
	if err != nil {
		return
	}

	// version check
	if !api.ignoreVersionCheck {
		var isLatest bool
		var latest string
		isLatest, latest, err = api.IsLatest()
		if err != nil {
			return
		}

		if !isLatest {
			notLatest := fmt.Sprintf(
				"You're not using the latest majsoul api!\n"+
					"Now using: %s, Latest: %s\n"+
					"Please upgrade this package.",
				ApiVersion,
				latest,
			)
			if api.panicOnVersionMismatch {
				panic(notLatest)
			} else {
				fmt.Println(notLatest)
			}
		}
	}
	return
}

// GetRegionUrl Get region url current api
func (a *MajsoulAPI) GetRegionUrl() string {
	return a.regions[a.Region]
}

// GetWebSocketUrl Pick up one websocket server from the slice
func (a *MajsoulAPI) GetWebSocketUrl() string {
	if a.activeWebSocketUrl == "" {
		rand.Seed(time.Now().Unix())
		n := rand.Int() % len(a.websocketServers)
		a.activeWebSocketUrl = a.websocketServers[n]
	}
	return "wss://" + a.activeWebSocketUrl
}
