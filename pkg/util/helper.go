package util

import (
	"crypto/md5"
	"encoding/hex"
	"net/mail"

	"github.com/google/uuid"
)

func GenerateSessionID() string {
	return uuid.NewString()
}

func ValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
