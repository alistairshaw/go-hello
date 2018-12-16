package mailsend

import (
	"go-hello/internal/services/config"
	"log"
	"net/smtp"
	"strconv"
)

//Send : send an email
func Send(subject string, message string) bool {
	c, configError := config.SMTP()
	if configError != nil {
		log.Print(configError.Error())
		return false
	}
	auth := smtp.PlainAuth(
		"",
		c.SMTPUser,
		c.SMTPPassword,
		c.SMTPHost)

	to := c.SMTPEmailAddress
	from := c.SMTPEmailAddress

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Hello there\n\n" +
		message

	host := c.SMTPHost + ":" + strconv.Itoa(c.SMTPPort)
	err := smtp.SendMail(
		host,
		auth,
		c.SMTPEmailAddress,
		[]string{c.SMTPEmailAddress},
		[]byte(msg))

	if err != nil {
		log.Print(err)
		return false
	}
	return true
}
