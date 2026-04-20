package main

import (
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name  string `form:"name" binding:"required"` // 绑定查询参数和表单数据，binding规则表示该字段是必填的
	Age   int    `form:"age" binding:"required"`
	Email string `form:"email" binding:"required,email"` // binding规则表示该字段必须是一个有效的邮箱地址
	Phone string `form:"phone" binding:"phone"`          // binding规则表示该字段必须通过自定义的phone验证
}

type UserUri struct {
	Name string `uri:"name" binding:"min=3"`           // 绑定uri路径参数，binding规则表示该字段长度必须大于等于3
	Age  int    `uri:"age" binding:"omitempty,gte=14"` // binding规则表示该字段为零值时不进行验证，否则必须大于等于14
}

type UserJson struct {
	Name string `json:"name"`
	Age  int    `json:"age" binding:"gte=14,lte=140"` // 绑定JSON请求体，binding规则表示年龄必须在14到140之间
}

type UserHeader struct {
	Name string `header:"X-Name"`
	Age  int    `header:"X-Age"`
}

// 自定义验证函数
func ValidatePhone(fl validator.FieldLevel) bool {
	reg := regexp.MustCompile(`^1[3-9]\d{9}$`)  // 简单的中国大陆手机号正则表达式
	return reg.MatchString(fl.Field().String()) // 验证字段的值是否匹配正则表达式
}

func main() {
	r := gin.Default()

	// 绑定查询参数
	r.GET("/query", func(c *gin.Context) {
		var user User
		if err := c.ShouldBind(&user); err != nil {
			c.String(400, "Error: %v", err)
			return
		}
		c.JSON(200, user)
	})

	// 绑定uri路径参数
	r.GET("/uri/:name/:age", func(c *gin.Context) {
		var user UserUri
		if err := c.ShouldBindUri(&user); err != nil {
			c.String(400, "Error: %v", err)
			return
		}
		c.JSON(200, user)
	})

	// 绑定JSON请求体
	r.POST("/json", func(c *gin.Context) {
		var user UserJson
		if err := c.ShouldBindJSON(&user); err != nil {
			c.String(400, "Error: %v", err)
			return
		}
		c.JSON(200, user)
	})

	// 绑定表单数据
	r.POST("/form", func(c *gin.Context) {
		var user User
		if err := c.ShouldBind(&user); err != nil {
			c.String(400, "Error: %v", err)
			return
		}
		c.JSON(200, user)
	})

	// 绑定header参数
	r.GET("/header", func(c *gin.Context) {
		var user UserHeader
		if err := c.ShouldBindHeader(&user); err != nil {
			c.String(400, "Error: %v", err)
			return
		}
		c.JSON(200, user)
	})

	// 自定义验证规则
	// 注册自定义验证函数，参数"phone"是绑定标签中使用的名称，ValidatePhone是验证函数
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("phone", ValidatePhone)
	}

	// 使用
	r.POST("/valid", func(c *gin.Context) {
		var user User
		if err := c.ShouldBind(&user); err != nil {
			c.String(400, "Error: %v", err)
			return
		}
		c.JSON(200, user)
	})

	r.Run(":8080")
}
