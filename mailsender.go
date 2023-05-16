package mailsender

import (
	"net/smtp"
)

func SendMailGoogle(mailFrom, mailTo, key string, err error) {
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

func SendMail(mailFrom, mailTo, key, host, addr, subject string, err error) {
	auth := smtp.PlainAuth(
		"",
		mailFrom,
		key,
		host,
	)

	msg := "Subject: " + subject + "\n" + err.Error()

	errSend := smtp.SendMail(
		addr,
		auth,
		mailFrom,
		[]string{mailTo},
		[]byte(msg),
	)

	if errSend != nil {
		return
	}
}
