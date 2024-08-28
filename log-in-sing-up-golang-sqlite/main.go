package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

var HashedPass string

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	r.GET("/sign-in", func(c *gin.Context) {
		c.HTML(200, "sing-in.html", gin.H{})
	})

	r.GET("/sign-up", func(c *gin.Context) {
		c.HTML(200, "sing-up.html", gin.H{})
	})

	r.POST("/sign-in", signIn)

	r.POST("/sign-up", func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")
		hashing(password)
		database(email, HashedPass)
		c.HTML(200, "sing-up.html", gin.H{})
	})

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func hashing(password string) {
	pass := []byte(password)

	// Hashing the password
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return
	}
	HashedPass = string(hash)
}

func database(email, hashedPass string) {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Create table if not exists
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, email TEXT, password TEXT)")
	if err != nil {
		log.Fatalf("Error preparing statement: %v", err)
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalf("Error executing statement: %v", err)
	}

	// Insert user into the database
	_, err = db.Exec("INSERT INTO users (email, password) VALUES (?, ?)", email, hashedPass)
	if err != nil {
		log.Fatalf("Error inserting user into database: %v", err)
	}
}

func signIn(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	hashing(password)
	err := saveData(email, HashedPass)
	if err != nil {
		log.Fatalf("Error inserting user into database: %v", err)
	}
	return

	c.HTML(http.StatusOK, "sing-in.html", gin.H{
		"Success": "Data saved succesfuly",
	})
}
