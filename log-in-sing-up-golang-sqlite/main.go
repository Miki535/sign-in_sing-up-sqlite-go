package main

import (
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

	r.POST("/sign-up", signUp)
	r.POST("/sign-in", singIn)

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

func signUp(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	hashing(password)
	err := saveData(email, HashedPass)
	if err != nil {
		log.Fatalf("Error inserting user into database: %v", err)
	}
	return

	c.HTML(http.StatusOK, "sing-up.html", gin.H{
		"Success": "Data saved succesfuly",
	})
}
func singIn(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	hashing(password)
	getData(email, HashedPass)
	if result == 1 {
		http.Redirect(c.Writer, c.Request, "/", 302)
	} else {
		http.Redirect(c.Writer, c.Request, "/", 302)
	}
	return
	c.HTML(http.StatusOK, "sing-in.html", gin.H{})
}
