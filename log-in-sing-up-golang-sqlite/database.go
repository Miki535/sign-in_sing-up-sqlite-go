package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB
var result int
var TokenCheckedResult bool

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	createTable()
}

func init() {
	var err error
	db, err = sql.Open("sqlite3", "./tokens.db")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	createTokenTable()
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
		fmt.Println(err)
	}
}

func resetPassword(email string, password string) {

}
func createTokenTable() {
	query := `
    CREATE TABLE IF NOT EXISTS tokens (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        tokens INTEGER,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    `
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Failed to create tokens table:", err)
	}
}

func saveToken() error {
	insertQuery := `INSERT INTO tokens (tokens, created_at) VALUES (?, datetime('now'))`
	_, err := db.Exec(insertQuery, Token)
	if err != nil {
		log.Println("Error saving token:", err)
		return err
	}

	deleteQuery := `DELETE FROM tokens WHERE created_at < datetime('now', '-2 minutes')`
	_, err = db.Exec(deleteQuery)
	if err != nil {
		log.Println("Error deleting old tokens:", err)
		return err
	}

	return nil
}
func IfTokenExists(token int) {
	query := `SELECT EXISTS(SELECT 1 FROM tokens WHERE tokens = ? LIMIT 1)`
	err := db.QueryRow(query, token).Scan(&TokenCheckedResult)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("Error checking token: %v", err)
	}
}
