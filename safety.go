package majsoul_api

import (
	"encoding/json"
	"net/http"
)

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

// Ignore version check when creating socket clients.
func (a *MajsoulAPI) IgnoreVersionCheck(ignore bool) {
	a.ignoreVersionCheck = ignore
}

// When ignoreVersionCheck is false, this controls whether panic(version_mismatch) works.
// Or just print out the warning information.
func (a *MajsoulAPI) PanicOnVersionMismatch(panic bool) {
	a.panicOnVersionMismatch = panic
}

// GetResourceVersion Get resource version now
func (a *MajsoulAPI) GetResourceVersion() string {
	version, _ := getVersion(a.Base)
	return version.Version
}

// Get code.js version now.
func (a *MajsoulAPI) GetCodeVersion() string {
	version, _ := getVersion(a.Base)
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

	res := resVersion{}
	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		return
	}

	prefix = res.Res
	return
}

func (a *MajsoulAPI) refreshResVersion() error {
	version, err := getVersion(a.Base)
	if err != nil {
		panic(err)
	}

	prefix, err := getResVersion(a.Base, version.Version)
	if err != nil {
		return err
	}

	a.resourceVersion = &prefix
	return nil
}

func (a *MajsoulAPI) getCachedResVersion() error {
	if a.resourceVersion == nil {
		return a.refreshResVersion()
	}
	return nil
}

// IsLatest Determine whether it's the latest lobby version.
// It returns a boolean and a detailed version string.
func (a *MajsoulAPI) IsLatest() (bool, string, error) {
	version, err := getVersion(a.Base)
	if err != nil {
		return false, ApiVersion, err
	}

	if version.Version == LatestTestedVersion {
		return true, ApiVersion, nil
	}

	err = a.getCachedResVersion()
	if err != nil {
		return false, version.Version, err
	}

	return (*a.resourceVersion)["res/proto/liqi.json"].Prefix == "v"+ApiVersion,
		(*a.resourceVersion)["res/proto/liqi.json"].Prefix,
		nil
}
