package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
this starts a server and run here: http://127.0.0.1:8080/hello
*/

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})

	r.Run(":8080")
}
