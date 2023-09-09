package util

import (
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
