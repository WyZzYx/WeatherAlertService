package utils

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
)

func SendEmail(to, body string) {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("MAIL_FROM"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Weather Notification")
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(
		"smtp.gmail.com",
		587,
		os.Getenv("MAIL_FROM"),
		os.Getenv("MAIL_PASSWORD"),
	)
	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println("Error sending email:", err)
	} else {
		fmt.Println("Email sent to:", to)
	}
}
