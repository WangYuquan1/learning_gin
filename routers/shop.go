package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func loadShopHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello loadShopHandler!",
	})
}

func LoadShop(e *gin.Engine) {
	e.GET("/loadShop", loadShopHandler)
}
