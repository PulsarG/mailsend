package mailsender

import (
	"net/smtp"
)

func SendMail(mailFrom, mailTo, key string, err error) {
	auth := smtp.PlainAuth(
		"",
		mailFrom,
		key,
		"smtp.gmail.com",
	)

	msg := "Subject: Enigma error" + "\n" + err.Error()

	errSend := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		mailFrom,
		[]string{mailTo},
		[]byte(msg),
	)

	if errSend != nil {
		return
	}
}
