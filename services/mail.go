package services

import (
	"net/smtp"

	"../beans"
	"../env"
)

//SendMail 寄發通知郵件
func SendMail(params *beans.SendMail) error {
	cfg := env.GetEnv()
	host := cfg.Mail.Host + ":" + cfg.Mail.Port
	auth := smtp.PlainAuth("", cfg.Mail.User, cfg.Mail.Password, cfg.Mail.Host)
	message := []byte(
		"Subject: " + params.GetSubject() + "\r\n" +
			"To: " + params.GetTo()[0] + "\r\n" +
			"From: " + cfg.Mail.User + "\r\n" +
			"Content-Type: text/plain; charset=UTF-8" + "\r\n" +
			"\r\n" +
			params.GetContent() + "\r\n" +
			"\r\n",
	)
	err := smtp.SendMail(host, auth, cfg.Mail.User, params.GetTo(), message)
	return err
}
