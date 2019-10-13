package main

import (
	"fmt"
	api "github.com/Yesterday17/majsoul_api"
	"github.com/Yesterday17/majsoul_api/lq"
	"time"
)

const (
	MajsoulBase   = "www.majsoul.com/1"
	MajsoulApi    = "api.majsoul.com:2501"
	MajsoulServer = "wss://mj-srv-5.majsoul.com:4101"
)

func main() {
	err := api.SendRegisterCode(MajsoulApi, "t@ab.cd")
	if err != nil {
		fmt.Println(err)
	}

	isLatest, latest := api.IsLatest(MajsoulBase)
	fmt.Println("isLatest", isLatest, latest)
	fmt.Println("resourceVersion", api.GetResourceVersion(MajsoulBase))
	fmt.Println("codeVersion", api.GetCodeVersion(MajsoulBase))

	//api.IgnoreVersionCheck(true)
	//api.PanicOnVersionMismatch(false)
	lobby, _ := api.NewLobbyClient(MajsoulBase, MajsoulServer)
	game, _ := api.NewGameClient(MajsoulBase, MajsoulServer)

	go lobby.Listen()

	// Login
	login := <-lobby.Login(&lq.ReqLogin{
		Account:   "jaymen.matthews@thtt.us",
		Password:  api.EncodePassword("jaymen.matthews@thtt"),
		Reconnect: false,
		Device: &lq.ClientDeviceInfo{
			DeviceType: "pc",
			Os:         "mac",
			OsVersion:  "",
			Browser:    "firefox",
		},
		RandomKey:      api.GenerateDeviceId(),
		ClientVersion:  api.GetResourceVersion(MajsoulBase),
		GenAccessToken: true,
		CurrencyPlatforms: []uint32{
			// 1, // Google Play
			2, // China
		},
		Type: 0,
	})
	fmt.Println(login)

	// after login
	serverTime := <-lobby.FetchServerTime()
	fmt.Println(serverTime.ServerTime)

	settings := <-lobby.FetchServerSettings()
	fmt.Println(settings)

	info := <-lobby.FetchConnectionInfo()
	fmt.Println(info.ClientEndpoint)

	lobby.AddListener("NotifyAccountUpdate", func(wrapper lq.Wrapper) {
		fmt.Println(wrapper)
	})

	lobby.AddListener("NotifyAnotherLogin", func(wrapper lq.Wrapper) {
		fmt.Println(wrapper)
	})

	lobby.AddListener("NotifyAccountLogout", func(wrapper lq.Wrapper) {
		fmt.Println(wrapper)
	})

	lobby.AddListener("NotifyClientMessage", func(wrapper lq.Wrapper) {
		fmt.Println(wrapper)
	})

	lobby.AddListener("NotifyServerSetting", func(wrapper lq.Wrapper) {
		fmt.Println(wrapper)
	})

	lobby.AddListener("NotifyVipLevelChange", func(wrapper lq.Wrapper) {
		fmt.Println(wrapper)
	})

	lastHeartBeatTime := time.Now().Second()
	go func() {
		time.AfterFunc(time.Minute*6, func() {
			// if game.LobbyNetMgr.Inst.isOK
			lobby.Heartbeat(&lq.ReqHeatBeat{
				NoOperationCounter: uint32(time.Now().Second() - lastHeartBeatTime),
			})
		})
	}()

	// client heartbeat
	go func() {
		time.AfterFunc(time.Second, func() {
			if time.Now().Second()-lastHeartBeatTime > 2400 {
				// if game.LobbyNetMgr.Inst.isOK
				lobby.Heartbeat(&lq.ReqHeatBeat{
					NoOperationCounter: 0,
				})
				lastHeartBeatTime = time.Now().Second()
			}
		})
	}()

	//go func() {
	//	time.AfterFunc(time.Second, func() {
	//		// set localStorage.dolllt to time.Now()
	//	})
	//}()

	announcement := <-lobby.FetchAnnouncement()
	for _, j := range announcement.Announcements {
		fmt.Printf("[公告] %s\n%s\n\n", j.Title, j.Content)
	}

	record := <-lobby.FetchGameRecord(&lq.ReqGameRecord{
		GameUuid: "191006-9d228ebb-8d8c-48c0-9c13-bed0c1fbfd24",
	})
	fmt.Println("牌谱: 191006-9d228ebb-8d8c-48c0-9c13-bed0c1fbfd24")
	players := map[uint32]string{}
	for _, id := range record.Head.Accounts {
		players[id.Seat] = id.Nickname
	}
	for _, p := range record.Head.Result.Players {
		fmt.Printf("玩家:%s 位置:%d 点数:%d\n", players[p.Seat], p.Seat, p.TotalPoint)
	}

	i := <-lobby.FetchAccountInfo(&lq.ReqAccountInfo{AccountId: 114514})
	fmt.Println(i.Room)

	si := <-lobby.FetchAccountStatisticInfo(&lq.ReqAccountStatisticInfo{AccountId: 114514})
	fmt.Printf("%+v", si)

	_ = game
	select {}
}
