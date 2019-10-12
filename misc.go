package majsoul_api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/uuid"
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

// TODO
func SendRegisterCode(email string) error {
	return nil
}
