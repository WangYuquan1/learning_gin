package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func blogaddHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello blogaddHandler",
	})
}

func BlogShop(e *gin.Engine) {
	e.GET("/add", blogaddHandler)
}
