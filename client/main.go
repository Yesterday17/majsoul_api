package main

import (
	"fmt"
	api "github.com/Yesterday17/majsoul_api"
	"github.com/Yesterday17/majsoul_api/lq"
)

const (
	MajsoulServer = "wss://mj-srv-5.majsoul.com:4101"
)

func main() {
	lobby, _ := api.NewLobbyClient(MajsoulServer)
	game, _ := api.NewGameClient(MajsoulServer)

	go lobby.Listen()

	(func() {
		ch := make(chan *lq.ResConnectionInfo)
		lobby.FetchConnectionInfo(ch)
		info := <-ch
		fmt.Println(info.ClientEndpoint)
	})()

	(func() {
		ch := make(chan *lq.ResServerTime)
		lobby.FetchServerTime(ch)
		info := <-ch
		fmt.Println(info.ServerTime)
	})()

	_ = game
}
