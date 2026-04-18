package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GETHandler(c *gin.Context) {
	fmt.Println(c.Request.Method, c.Request.URL.String()) // 输出请求的类型和 URL 路径

	// 使用 Gin 的 JSON 方法返回一个 JSON 响应，包含状态和消息字段
	// JSON 方法封装了设置响应头和编码 JSON 数据的过程，使代码更简洁易读
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "GET request received",
	})
}

func POSTHandler(c *gin.Context) {
	fmt.Println(c.Request.Method, c.Request.URL.String()) // 输出请求的类型和 URL 路径

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "POST request received",
	})
}

func main() {
	// 创建一个默认的 Gin 引擎实例
	r := gin.Default()

	//
	r.StaticFile("/favicon.ico", "./favicon.ico")

	// 定义 GET 和 POST 请求的路由和处理函数，当资源路径为 "/get" 时，调用 GETHandler 处理 GET 请求
	r.GET("/get", GETHandler)
	r.POST("/post", POSTHandler)

	// 启动 Gin 服务器，监听在默认的 8080 端口
	fmt.Println("Gin Server is already running on http://localhost:8080")
	r.Run(":8080")
}
