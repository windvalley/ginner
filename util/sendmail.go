package util

import (
	"gopkg.in/gomail.v2"

	"use-gin/config"
)

// SendMail using for send a alert mail to manager.
//    e.g.:
//mailTo := []string{
//	"manager@sre.im",
//}
//subject := "[alert] use-gin error"
//body := "some html codes"
//if err := util.SendMail(mailTo, subject, body); err != nil {
//return err
//}
func SendMail(mailTo []string, subject string, body string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(config.Conf().Mail.User, "alias name"))
	m.SetHeader("To", mailTo...)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer(
		config.Conf().Mail.SMTPHost,
		config.Conf().Mail.Port,
		config.Conf().Mail.User,
		config.Conf().Mail.Password,
	)

	err := d.DialAndSend(m)
	return err
}
