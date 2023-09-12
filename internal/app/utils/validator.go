package utils

import (
	"net/mail"
	"strings"
)

func ContainsWhiteSpace(password string) bool {
	return strings.Contains(password, " ")
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}