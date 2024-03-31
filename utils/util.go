package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func PasswordEncrypt(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	sum := h.Sum(nil)
	return hex.EncodeToString(sum)
}
