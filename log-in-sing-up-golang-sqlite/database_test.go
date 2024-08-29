package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestHol(t *testing.T) {
	email := "sherbantaras535@gmail.com"
	password := "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
	getData(email, password)
	fmt.Println(result)
}
