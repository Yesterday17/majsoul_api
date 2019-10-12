package main

import (
	"fmt"
	api "github.com/Yesterday17/majsoul_api"
)

const (
	MajsoulServer = "wss://mj-srv-5.majsoul.com:4101"
)

func main() {
	lobby, _ := api.NewLobbyClient(MajsoulServer)
	game, _ := api.NewGameClient(MajsoulServer)

	go lobby.Listen()

	info := <-lobby.FetchConnectionInfo()
	fmt.Println(info.ClientEndpoint)

	time := <-lobby.FetchServerTime()
	fmt.Println(time.ServerTime)

	_ = game
}
