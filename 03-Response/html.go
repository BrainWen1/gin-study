package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 加载html文件
	r.LoadHTMLGlob("template/*")

	// 挂载路由
	r.GET("/index", func(c *gin.Context) {
		// 渲染html文件
		c.HTML(200, "index.html", gin.H{
			"title":   "Go-Gin", // 在这里可以传递数据到html文件中，在html文件中可以使用 {{.title}} 来获取这个值
			"message": "Hello, Go-Gin!",
		})
	})

	r.Run(":8080")
}
