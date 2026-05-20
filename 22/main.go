package main

import (
	"net/smtp"

	"github.com/jhillyerd/enmime"
)

func main() {
	smtpHost := "my-mail-server:25"
	smtpAuth := smtp.PlainAuth(
		"example.com",
		"example-user",
		"example-password",
		"auth.example.com")

	sender := enmime.NewSMTP(smtpHost,smtpAuth)

	master := enmime.Builder().

	From()
}