package main

import (
	"gin-study/03-Response/res"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		res.Ok(c, "hello", gin.H{})
	})
	r.GET("/login", func(c *gin.Context) {
		res.Ok(c, "login", gin.H{})
	})
	r.GET("/fail", func(c *gin.Context) {
		res.Fail(c, 1002, "user not found", nil)
	})

	r.POST("/fail", func(c *gin.Context) {
		res.FailWithData(c, 1003, nil)
	})

	r.Run(":8080")
}
