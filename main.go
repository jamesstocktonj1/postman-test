package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	VERSION = "1.0.0"
)

func main() {
	router := gin.Default()

	router.GET("/ping", Ping)
	router.GET("/version", Version)

	router.Run(":8080")
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

func Version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"version": VERSION})
}
