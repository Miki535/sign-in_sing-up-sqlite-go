package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
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

	r.POST("/sign-in", func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")
		hashing(password)
		database("./data.db", email, HashedPass)
		c.HTML(200, "sing-in.html", gin.H{})
	})

	r.POST("/sign-up", func(c *gin.Context) {
		c.HTML(200, "sing-up.html", gin.H{})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func hashing(password string) {
	pass := []byte(password)

	// Hashing the password
	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
		log.Println("Error hashing password")
	}
	HashedPass = string(hash)
}
