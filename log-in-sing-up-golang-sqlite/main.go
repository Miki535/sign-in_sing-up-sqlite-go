package main

import (
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"strconv"
)

var HashedPass string
var email string

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
	r.GET("/resetPass", func(c *gin.Context) {
		token := c.Query("token")
		tokenToINT, _ := strconv.Atoi(token)
		IfTokenExists(tokenToINT)
		if TokenCheckedResult == true {
			c.Redirect(http.StatusMovedPermanently, "/sign-up")
		} else {
			c.Redirect(http.StatusMovedPermanently, "/404")
		}

		if token != fmt.Sprint(Token) {
			c.Redirect(302, "/404")
		}

		c.HTML(200, "resetPass.html", gin.H{
			"token": token,
		})
	})

	r.POST("/resetPass", resetPass)

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
	password := c.PostForm("password")
	hashing(password)
	err := saveData(HashedPass, email)
	if err != nil {
		fmt.Errorf("Error saving data: %v", err)
	}
	Tokenizator()
	AlertOnEmail(email)
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

func resetPass(c *gin.Context) {
	newPassword := c.PostForm("newPassword")
	hashing(newPassword)

	c.HTML(200, "resetPass.html", gin.H{})
}
