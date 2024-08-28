package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB
var result int

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	} else {
		createTable()
	}

}

func createTable() {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        password TEXT,
        email TEXT
    );
    `
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
}

func saveData(password, email string) error {
	query := `INSERT INTO users (password, email) VALUES (?, ?)`
	_, err := db.Exec(query, password, email)
	return err
}

func getData(email string, password string) {
	err := db.QueryRow("SELECT 1 FROM users WHERE email = ? AND password = ?", email, password).Scan(&result)
	if err != nil {
		log.Fatal("Failed to get data:", err)
	}
}
