package res

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int
	Msg  string
	Data any
}

func response(c *gin.Context, code int, msg string, data any) {
	c.JSON(200, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// ok
func Ok(c *gin.Context, msg string, data any) {
	response(c, 0, msg, data)
}

func OkWithMsg(c *gin.Context, msg string) {
	response(c, 0, msg, nil)
}

func OkWithData(c *gin.Context, data any) {
	response(c, 0, "success", data)
}

// fail
var CodeMap = map[int]string{
	1001: "invalid parameter",
	1002: "user not found",
	1003: "internal server error",
}

func Fail(c *gin.Context, code int, msg string, data any) {
	response(c, code, msg, data)
}

func FailWithMsg(c *gin.Context, msg string) {
	response(c, 1001, msg, nil)
}

func FailWithData(c *gin.Context, code int, data any) {
	myCode, ok := CodeMap[code]
	if !ok {
		myCode = "unknown error"
	}

	response(c, code, myCode, data)
}
