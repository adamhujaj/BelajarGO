package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Endpoint for /ping
	r.GET("/ping", func(c *gin.Context) {
		c.HTML(http.StatusOK, "ping.html", nil)
	})

	// Endpoint for /rasul
	r.GET("/rasul", func(c *gin.Context) {
		c.HTML(http.StatusOK, "rasul.html", nil)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
