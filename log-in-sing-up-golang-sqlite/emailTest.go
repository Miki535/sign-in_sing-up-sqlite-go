package main

import (
	"fmt"
	"log"
	"net/smtp"
	"net/url"
)

func AlertOnEmail(email string) {
	secretKey := "uane nfkg zomc dtii"
	ownEmail := "sherbantaras535@gmail.com"

	auth := smtp.PlainAuth("", ownEmail, secretKey, "smtp.gmail.com")
	Tokenizator()
	confirmationURL := fmt.Sprintf("http://localhost:8080/resetPass?token=%s", url.QueryEscape(fmt.Sprint(Token)))
	subject := "Confirm Your Email Address"
	body := fmt.Sprintf("Click the following link to confirm your email address: <a href='%s'>Confirm Email</a>", confirmationURL)
	msg := []byte("Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail("smtp.gmail.com:587", auth, ownEmail, []string{email}, msg)
	if err != nil {
		log.Fatal(err)
	}

}
