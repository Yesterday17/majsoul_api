package majsoul_api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	majsoulSecret = "lailai"
)

func EncodePassword(password string) string {
	h := hmac.New(sha256.New, []byte(majsoulSecret))
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}

func GenerateDeviceId() string {
	return uuid.New().String()
}

// TODO
func IsLatest() bool {
	// 1. https://www.majsoul.com/1/version.json
	// 2. https://www.majsoul.com/1/resversion0.6.73.w.json
	// 3. res/proto/liqi.json -> prefix
	return true
}

func SendRegisterCode(api, email string) error {
	resp, err := http.Post(
		"https://"+api+"/api/user/sign_up_code",
		"application/x-www-form-urlencoded",
		strings.NewReader("type=email&email="+email))
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if string(body) != "{}" {
		return fmt.Errorf(string(body))
	}
	return nil
}
