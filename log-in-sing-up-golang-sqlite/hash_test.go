package main

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"testing"
)

var hashedPass string

func Hashing(password string) {
	pass := []byte(password)

	// Hashing the password
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
		log.Println("Error hashing password")
	}
	hashedPass = string(hash)
}

func TestToHashingFunc(t *testing.T) {
	var password string
	password = "test"
	Hashing(password)
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
	if err != nil {
		t.Errorf("Password does not match")
	} else {
		t.Log("Password match")
	}
}
