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

// EncodePassword Hash your clear password using the way majsoul uses
func (a *MajsoulAPI) EncodePassword(password string) string {
	h := hmac.New(sha256.New, []byte(a.Secret))
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}

// GenerateDeviceId Generate a uuid which majsoul uses as device id
func (a *MajsoulAPI) GenerateDeviceId() string {
	return uuid.New().String()
}

// SendRegisterCode Send an email to the email you provide and determine whether it succeeds
func (a *MajsoulAPI) SendRegisterCode(email string) error {
	resp, err := http.Post(
		a.systemEmailUrl+"/api/user/sign_up_code",
		"application/x-www-form-urlencoded",
		strings.NewReader("type=email&email="+email))
	if err != nil {
		return err
	}

	body, err := readResponse(resp)
	if err != nil {
		return err
	}

	if body != "{}" {
		return fmt.Errorf(string(body))
	}
	return nil
}

func readResponse(resp *http.Response) (string, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return string(body), nil
}
