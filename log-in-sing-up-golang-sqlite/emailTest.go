package main

import (
	"fmt"
	"log"
	"net/smtp"
)

var TestCode int

func SendTestCode(email string) {
	secretKey := "secretkey"
	own_email := "ownemail"

	auth := smtp.PlainAuth("", own_email, secretKey, "smtp.gmail.com")

	to := []string{email}
	Randomizer(101, 10001)
	msg := TestCode
	// send message on email
	err := smtp.SendMail("smtp.gmail.com:587", auth, own_email, to, []byte(fmt.Sprint(msg)))
	if err != nil {
		log.Fatal(err)
	}
}
