package main

import (
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

//SendMailToAddress sends mail
func SendMailToAddress(address, path, title, user, pw string) {
	mail := email.NewEmail()
	mail.From = "XKCD Subscriber"
	mail.To = []string{address}
	mail.Subject = "Todays XKCD comic: " + title
	mail.AttachFile(path)
	mail.Text = []byte("Todays comic is attached to this mail")
	mail.Send("mail.gmx.net:587", smtp.PlainAuth("", user, pw, "mail.gmx.net"))
	log.Println("Email sent to", address)
}
