package main

import (
	"fmt"
	api "github.com/Yesterday17/majsoul_api"
	"github.com/Yesterday17/majsoul_api/lq"
)

func main() {
	lobby := &api.Lobby{}
	game := &api.Game{}

	_ = lobby.Init()
	_ = game.Init()

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
}
