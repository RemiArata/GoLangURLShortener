package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

/*
this starts a server and run here: http://127.0.0.1:8080/hello
*/

var URLMap = make(map[string]string)

func main() {
	rand.Seed(time.Now().UnixNano())
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})

	r.GET("/shorten/:url", shortenURL)
	r.GET("/:shortURL", goToShortUrl)

	r.Run(":8080")
}

func shortenURL(c *gin.Context) {
	longURL := c.Param("url")
	shortURL := generateShortURL()
	URLMap[shortURL] = longURL
	fmt.Println(URLMap)
	c.JSON(http.StatusOK, gin.H{"shortURL": shortURL})
}

func generateShortURL() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const size = len(charset)
	short := make([]byte, 6)
	for i := range short {
		short[i] = charset[rand.Intn(size)]
	}
	return string(short)
}

func goToShortUrl(c *gin.Context) {
	shortURL := c.Param("shortURL")
	fmt.Println("the short url ", shortURL)
	longURL, ok := URLMap[shortURL]
	fmt.Println(ok)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"message": "short URL doesn't have corrosponding long URL"})
		return
	}
	fmt.Println(longURL)
	c.Redirect(http.StatusMovedPermanently, longURL)
}
