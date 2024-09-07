package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

var HashedPass string
var email string
var token string
var paSS string

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
	r.GET("/404", func(c *gin.Context) {
		c.HTML(200, "404.html", gin.H{})
	})
	r.GET("/confirm", func(c *gin.Context) {
		token = c.Query("token")
		if token == "" {
			c.JSON(400, gin.H{"error": "Token is required"})
			return
		}
		if token == paSS {
			c.Redirect(302, "/")
		} else {
			c.JSON(400, gin.H{"error": "Invalid token"})
		}

		c.HTML(200, "confirm.html", gin.H{
			"token": token,
		})
	})

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func hashing(password string) {
	pass := []byte(password)

	hash := sha256.Sum256(pass)
	HashedPass = fmt.Sprintf("%x", hash[:])
}

func signUp(c *gin.Context) {
	email = c.PostForm("email")
	paSS = c.PostForm("password")
	if email != "" {
		SendTestCode(email, paSS)
	}
	c.HTML(http.StatusOK, "sing-up.html", gin.H{})
}
func singIn(c *gin.Context) {
	email = c.PostForm("email")
	password := c.PostForm("password")
	hashing(password)
	getData(email, HashedPass)
	if result == 1 {
		http.Redirect(c.Writer, c.Request, "/", 302)
	} else {
		http.Redirect(c.Writer, c.Request, "/sign-up", 302)
	}
	return
	c.HTML(http.StatusOK, "sing-in.html", gin.H{})
}
