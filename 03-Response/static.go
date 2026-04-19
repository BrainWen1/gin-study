package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 将sta目录下的文件映射到/static路径下，访问/static/favicon.ico就会访问sta目录下的favicon.ico文件
	r.Static("/static", "./sta")

	r.Run(":8080")
}
