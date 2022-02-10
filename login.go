package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Login struct {
	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

/*　json parser
検証URL
curl --location --request POST 'http://localhost:8080/loginJSON' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user":"root",
    "password":"admin"
}'
*/
func loginJSON(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if json.User != "root" || json.Pssword != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}
