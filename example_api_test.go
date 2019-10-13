package majsoul_api

import (
	"fmt"
	"github.com/Yesterday17/majsoul_api/lq"
	"time"
)

func Example() {
	api, err := NewMajsoulAPI(MajsoulCN)
	if err != nil {
		fmt.Println(err)
	}

	err = api.SendRegisterCode("t@ab.ef")
	if err != nil {
		fmt.Println(err)
	}

	isLatest, latest, err := api.IsLatest()
	fmt.Println("isLatest", isLatest, latest)
	fmt.Println("resourceVersion", api.GetResourceVersion())
	fmt.Println("codeVersion", api.GetCodeVersion())

	//api.IgnoreVersionCheck(true)
	//api.PanicOnVersionMismatch(false)
	lobby, err := api.LobbyClient()
	if err != nil {
		panic(err)
	}

	game, err := api.GameClient()
	if err != nil {
		panic(err)
	}

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
		ClientVersion:  api.GetResourceVersion(),
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

func ExampleNewMajsoulAPI() {
	api, err := NewMajsoulAPI(MajsoulJP)
	if err != nil {
		fmt.Println(err)
	}

	// More operations here...
	_ = api
}
