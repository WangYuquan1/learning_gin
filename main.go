package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// router := gin.Default()
	// router.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "welcome to Learning Gin")
	// })
	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// router.GET("/user/:name/*action", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	action := c.Param("action")
	// 	action = strings.Trim(action, "/")
	// 	c.String(http.StatusOK, name+" is "+action) // http://localhost:8080/user/wang/drinking
	// })

	// router.GET("/user", func(c *gin.Context) {
	// 	name := c.DefaultQuery("name", "gest")
	// 	c.String(http.StatusOK, fmt.Sprintf("hello %s", name)) //http://localhost:8080/user?name=wang
	// })

	// router.POST("/form", func(c *gin.Context) {
	// 	types := c.DefaultPostForm("type", "post")
	// 	username := c.PostForm("username")
	// 	password := c.PostForm("password")
	// 	c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	// })

	// /* ルーターグループ
	// 検証URL：
	// curl http://localhost:8080/v1/login?name=wang
	// curl http://localhost:8080/v1/submit -X POST
	// curl http://localhost:8080/v2/submit -X POST
	// */
	// v1 := router.Group("/v1")
	// {
	// 	v1.GET("/login", login)
	// 	v1.GET("/submit", submit)
	// }

	// v2 := router.Group("/v2")
	// {
	// 	v2.POST("/login", login)
	// 	v2.POST("/submit", submit)
	// }

	// router.Run() // listen and serve on 0.0.0.0:8080 (for browser "localhost:8080")
	// }

	r := gin.Default()
	r.POST("loginJSON", loginJSON)
	if err := r.Run(); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
