package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "welcome to Learning Gin")
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, name+" is "+action) // http://localhost:8080/user/wang/drinking
	})

	router.Run() // listen and serve on 0.0.0.0:8080 (for browser "localhost:8080")
}
