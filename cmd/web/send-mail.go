package main

import (
	"log"
	"time"

	"github.com/IbrahimMohammedi/Bookings/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
)

func listenForMail() {

	go func() {
		for {
			msg := <-app.MailChain
			sendMsg(msg)
		}
	}()
}

func sendMsg(m models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		errorLog.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	email.SetBody(mail.TextHTML, string(m.Conetnt))

	err = email.Send(client)
	if err != nil {
		errorLog.Println(err)
	} else {
		log.Println("Email sent!")
	}
}
