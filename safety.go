package majsoul_api

import (
	"encoding/json"
	"net/http"
)

const (
	ApiVersion = "v0.6.58.w"
)

var ForceUsingOldApi = false

type majsoulVersion struct {
	Code    string `json:"code"`
	Version string `json:"version"`
}

type resVersion struct {
	Res map[string]prefixVersion `json:"res"`
}

type prefixVersion struct {
	Prefix string `json:"prefix"`
}

func GetResourceVersion(base string) string {
	version, _ := getVersion(base)
	return version.Version
}

func GetCodeVersion(base string) string {
	version, _ := getVersion(base)
	return version.Code
}

func getVersion(base string) (version majsoulVersion, err error) {
	var resp *http.Response
	var data string

	resp, err = http.Get("https://" + base + "/version.json")
	if err != nil {
		return
	}

	data, err = readResponse(resp)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(data), &version)
	return
}

func getResVersion(base, version string) (prefix map[string]prefixVersion, err error) {
	var resp *http.Response
	var data string

	resp, err = http.Get("https://" + base + "/resversion" + version + ".json")
	if err != nil {
		return
	}

	data, err = readResponse(resp)
	if err != nil {
		return
	}

	resv := resVersion{}
	err = json.Unmarshal([]byte(data), &resv)
	if err != nil {
		return
	}

	prefix = resv.Res
	return
}

func IsLatest(base string) (bool, string) {
	version, err := getVersion(base)
	if err != nil {
		panic(err)
	}

	prefix, err := getResVersion(base, version.Version)
	if err != nil {
		panic(err)
	}

	return prefix["res/proto/liqi.json"].Prefix == ApiVersion, prefix["res/proto/liqi.json"].Prefix
}