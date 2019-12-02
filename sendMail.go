package main

import (
	"log"
	"net/smtp"
	"sync"

	"github.com/jordan-wright/email"
)

//SendMailToAddress sends mail and signals a waitgroup that it is done
func SendMailToAddress(address, path, title, user, pw, smtpWithPort, smtpAddress string, waitgroup *sync.WaitGroup) {
	defer waitgroup.Done()
	mail := email.NewEmail()
	mail.From = "XKCD Subscriber"
	mail.To = []string{address}
	mail.Subject = "New XKCD comic: " + title
	mail.AttachFile(path)
	mail.Text = []byte("The new comic is attached to this mail")
	mail.Send(smtpWithPort, smtp.PlainAuth("", user, pw, smtpAddress))
	log.Println("Email sent to", address)
}
