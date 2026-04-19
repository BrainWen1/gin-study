package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("sta/*") // 加载HTML模板文件，sta目录下的所有文件都可以被加载

	// 展示文件
	r.GET("/show", func(c *gin.Context) {
		c.File("sta/favicon.ico")
	})
	// 下载文件1：直接使用FileAttachment方法，设置下载文件的路径和下载后的文件名，FileAttachment方法会自动设置响应头来告诉浏览器这是一个下载文件
	r.GET("/download", func(c *gin.Context) {
		c.FileAttachment("./sta/favicon.ico", "this.ico") // 下载文件，第一个参数是文件路径，第二个参数是下载后的文件名
	})
	// 下载文件2：手动设置响应头来告诉浏览器这是一个下载文件，然后使用File方法来发送文件
	r.GET("/download2", func(c *gin.Context) {
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename=this.ico")
		c.File("./sta/favicon.ico") // 下载文件，设置响应头来告诉浏览器这是一个下载文件
	})
	// 下载文件3：设置自定义响应头来告诉浏览器这是一个下载文件，然后使用File方法来发送文件
	// 但是这种方式需要前端在接收响应时根据自定义的响应头来处理下载文件
	r.GET("/download3", func(c *gin.Context) {
		c.Header("filename", "this.ico")
		c.Header("msg", "下载成功")
		c.File("./sta/favicon.ico")
	})
	r.GET("/download4", func(c *gin.Context) {
		c.HTML(200, "download.html", gin.H{
			"title":            "DOWNLOAD FILE",
			"downloadPath":     "/download2",
			"downloadFileName": "this.html",
		})
	})

	r.Run(":8080")
}
