package majsoul_api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type majsoulConfig struct {
	GoodsSheleveId string      `json:"goods_sheleve_id"`
	Ip             []majsoulIP `json:"ip"`
	SystemEmailUrl string      `json:"system_email_url"`
}

type majsoulIP struct {
	Name       string            `json:"name"`
	RegionUrls map[string]string `json:"region_urls"`
}

type majsoulWebsocketServer struct {
	Servers []string `json:"servers"`
}

func (a *MajsoulAPI) initConfig() error {
	err := a.getCachedResVersion()
	if err != nil {
		return err
	}

	res, err := http.Get("https://" + a.Base + "/" + (*a.resourceVersion)["config.json"].Prefix + "/config.json")
	if err != nil {
		return err
	}

	data, err := readResponse(res)
	if err != nil {
		return err
	}

	config := &majsoulConfig{}
	err = json.Unmarshal([]byte(data), &config)
	if err != nil {
		return err
	}

	a.systemEmailUrl = config.SystemEmailUrl
	a.regions = config.Ip[0].RegionUrls
	return nil
}

func (a *MajsoulAPI) initWebSocket() error {
	res, err := http.Get(a.GetRegionUrl() + "?service=ws-gateway&protocol=ws&ssl=true")
	if err != nil {
		return err
	}

	data, err := readResponse(res)
	if err != nil {
		return err
	}

	config := &majsoulWebsocketServer{}
	err = json.Unmarshal([]byte(data), &config)
	if err != nil {
		return err
	}

	if len(config.Servers) < 1 {
		return fmt.Errorf("empty websocket server list")
	}
	a.websocketServers = config.Servers
	return nil
}
