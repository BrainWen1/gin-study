package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func init() {
	// 创建翻译器
	uni := ut.New(zh.New())            // 获取中文翻译器
	trans, _ = uni.GetTranslator("zh") // 获取中文翻译器实例

	// 注册翻译器
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		_ = zh_translations.RegisterDefaultTranslations(v, trans)
	}
}

func ValidateErr(err error) string {
	// 将验证错误转换为翻译后的字符串
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error()
	}

	// 将每个验证错误翻译成中文，并拼接成一个字符串返回
	var list []string
	for _, e := range errs {
		list = append(list, e.Translate(trans))
	}
	return strings.Join(list, ";")
}

type User struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func main() {
	r := gin.Default()

	// 注册路由
	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			// 参数验证失败
			c.String(400, ValidateErr(err))
			return
		}

		// 参数验证成功
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Hello, %s! Your email is %s", user.Name, user.Email),
		})
	})

	// 启动HTTP服务器
	r.Run()
}
