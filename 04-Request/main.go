package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 查询参数 http://127.0.0.1:8080/query?name=zerotwo&list=123&list=iloveu
	r.GET("/query", func(c *gin.Context) {
		// 三种获取查询参数的方法
		name := c.Query("name")
		age := c.DefaultQuery("age", "18") // 如果age参数不存在，则默认值为18
		list := c.QueryArray("list")

		fmt.Println(name, age, list)
		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
			"list": list,
		})
	})

	// URI路径参数 http://127.0.0.1:8080/user/zerotwo/20
	r.GET("/user/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")

		fmt.Println(name, age)
		c.JSON(200, gin.H{
			"name": name,
			"age":  age,
		})
	})

	// 表单参数 http://127.0.0.1:8080/form
	r.POST("/form", func(c *gin.Context) {
		name := c.PostForm("name")
		age := c.DefaultPostForm("age", "18")
		gender, ok := c.GetPostForm("gender")

		fmt.Println(name, age, gender, ok)
		c.JSON(200, gin.H{
			"name":         name,
			"age":          age,
			"gender":       gender,
			"genderExists": ok,
		})
	})

	// 表单传文件 http://127.0.0.1:8080/upload
	r.POST("/upload", func(c *gin.Context) {
		// 也可以正常获取表单参数
		name := c.PostForm("name")

		// 获取上传的文件
		fileHeader, err := c.FormFile("file")
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		// 保存上传的文件到服务器
		err = c.SaveUploadedFile(fileHeader, "./"+fileHeader.Filename)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(name, fileHeader.Filename)
		c.JSON(200, gin.H{
			"name":     name,
			"filename": fileHeader.Filename,
			"msg":      "文件上传成功",
		})
	})

	// 多文件上传
	r.POST("/multi-upload", func(c *gin.Context) {
		name := c.PostForm("name")

		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		files := form.File["files"] // 获取多个文件
		if len(files) == 0 {
			c.JSON(400, gin.H{"error": "无文件上传"})
			return
		}
		var filenames []string

		// 遍历文件列表，保存每个文件到服务器本地
		for _, fileHeader := range files {
			err := c.SaveUploadedFile(fileHeader, "./UploadedFiles/"+fileHeader.Filename)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			filenames = append(filenames, fileHeader.Filename)
		}
		c.JSON(200, gin.H{
			"name":      name,
			"filenames": filenames,
			"msg":       "文件上传成功",
		})
		fmt.Println(name, filenames)
	})

	// body阅后即焚的处理
	r.POST("/body", func(c *gin.Context) {
		// 第一次读取body, GetRawData会将请求体读取到内存中，并且请求体会被清空
		bodyBytes, err := c.GetRawData()
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(bodyBytes)

		// 处理办法
		// 将读取到的bodyBytes重新放回请求体中，这样后续的读取就不会得到空了
		c.Request.Body = io.NopCloser(bytes.NewReader(bodyBytes))

		// 再次读取body会得到空，因为GetRawData已经读取了请求体
		name, ok := c.GetPostForm("name")
		age := c.PostForm("age")
		gender := c.PostForm("gender")
		fmt.Println(name, age, gender)

		c.JSON(200, gin.H{
			"body":   string(bodyBytes),
			"name":   name,
			"age":    age,
			"gender": gender,
			"ok":     ok,
		})
	})

	r.Run(":8080")
}
