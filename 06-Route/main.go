// RESTful 是行业通用的接口设计规范，核心是用[HTTP 方法]表示对资源的操作，而不是把操作写在 URL 里
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func listUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "List of users",
	})
}

func createUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "User created",
	})
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "User details",
		"id":      id,
	})
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "User updated",
		"id":      id,
	})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": "User deleted",
		"id":      id,
	})
}

func LoggerA() gin.HandlerFunc {
	return func(c *gin.Context) {
		// A 前置
		fmt.Println("A 前置")

		c.Next() // 处理请求

		// A 后置
		fmt.Println("A 后置")
	}
}

func LoggerB() gin.HandlerFunc {
	return func(c *gin.Context) {
		// B 前置
		fmt.Println("B 前置")

		c.Next() // 处理请求

		// B 后置
		fmt.Println("B 后置")
	}
}

func Authenticate(c *gin.Context) {
	fmt.Println("认证完成")
}

func Login(c *gin.Context) {
	fmt.Println("登录完成")
}

func main() {
	r := gin.Default()

	// 全局中间件 : A 前置 -> B 前置 -> 处理请求 -> 请求处理完成 -> B 后置 -> A 后置
	r.Use(LoggerA(), LoggerB())

	// 基础路由
	// 局部中间件：单路由
	r.GET("/ping", Authenticate, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "认证成功",
		})
	})

	// 路由分组
	// 局部中间件：路由分组
	v1 := r.Group("/api/v1", Authenticate, Login)
	{
		userGroup := v1.Group("/users")
		{
			userGroup.GET("", listUsers)         // 显示用户列表
			userGroup.POST("", createUser)       // 创建新用户
			userGroup.GET("/:id", getUser)       // 获取用户详情
			userGroup.PUT("/:id", updateUser)    // 更新用户信息
			userGroup.DELETE("/:id", deleteUser) // 删除用户
			userGroup.GET("/test", func(c *gin.Context) {
				fmt.Println("处理请求")
				c.JSON(200, gin.H{
					"message": "Test route",
				})
				fmt.Println("请求处理完成")
			})
		}
	}

	r.Run(":8080")
}
