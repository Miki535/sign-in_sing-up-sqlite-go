package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/sing-in", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sing-in.html", gin.H{})
	})

	r.GET("/sing-up", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sing-up.html", gin.H{})
	})

	r.POST("/sing-in", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sing-in.html", gin.H{})
	})

	r.POST("/sing-up", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sing-up.html", gin.H{})
	})

	return r
}

func TestPostSingIn(t *testing.T) {
	router := setupRouter()

	// Створення POST-запиту
	req, _ := http.NewRequest("POST", "/sing-in", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Перевірка статус-коду відповіді
	assert.Equal(t, http.StatusOK, w.Code)

	// Перевірка вмісту HTML-сторінки
	assert.Contains(t, w.Body.String(), "Sign In")
}

func TestPostSingUp(t *testing.T) {
	router := setupRouter()

	//Створення POST-запиту
	req, _ := http.NewRequest("POST", "/sign-up", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	//Перевірка статус-коду відповіді
	assert.Equal(t, http.StatusOK, w.Code)

	//Перевірка вмісту HTML-сторінки
	assert.Contains(t, w.Body.String(), "Sign Up")
}

func TestGetHomePage(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Contains(t, w.Body.String(), "Hello World!")
}

func TestGetSingIn(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/sign-in", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Contains(t, w.Body.String(), "Sign In")
}

func TestGetSingUp(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/sign-up", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	assert.Contains(t, w.Body.String(), "Sign Up")
}
