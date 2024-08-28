package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"testing"
)

func TestHol(t *testing.T) {
	email := "sherbantaras535@gmail.com"
	password := "test"
	getData(email, password)
	fmt.Println(result)
}
