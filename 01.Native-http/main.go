// Gin 解决了什么原生http服务的痛点？

// 1.参数解析与验证：Gin 提供了方便的参数解析和验证功能，可以轻松地从请求中提取参数并进行验证。原生 HTTP 需要手动解析和验证参数，代码较为繁琐。

// 2.路由管理：Gin 提供了强大的路由管理功能，可以轻松地定义路由和处理函数。原生 HTTP 需要手动管理路由，代码较为复杂。

// 3.响应处理：Gin 提供了方便的响应处理功能，可以轻松地返回 JSON、HTML 等格式的响应。原生 HTTP 需要手动构造响应，代码较为繁琐。

// 4.中间件支持：Gin 支持中间件，可以在请求处理过程中插入自定义逻辑。原生 HTTP 需要手动实现中间件功能，代码较为复杂。

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// response 结构体定义了 HTTP 响应的格式，包括状态码、消息和数据字段
type response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// GET 函数处理 GET 请求，接收 HTTP 请求并返回 JSON 格式的响应
func GET(res http.ResponseWriter, rep *http.Request) {
	fmt.Println(rep.Method, rep.URL.String()) // 输出请求的类型和 URL 路径

	// 构造响应数据并将其编码为 JSON 格式
	data, err := json.Marshal(response{
		Status:  200,                                   // HTTP 状态码，表示请求成功
		Message: "GET request received",                // 响应消息，说明请求已被接收
		Data:    "This is a response to a GET request", // 响应数据，提供对 GET 请求的具体响应内容
	})
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	// 设置响应头，指定内容类型为 JSON，并将编码后的数据写入响应体
	res.Header().Set("Content-Type", "application/json")
	// 将 JSON 数据写入响应体，发送给客户端
	res.Write(data)
}

// POST 函数处理 POST 请求，接收 HTTP 请求并返回 JSON 格式的响应
func POST(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method, req.URL.String()) // 输出请求的类型和 URL 路径

	// 拿到 POST 请求的 body 内容，并将其读取到变量 body 中
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError) // 如果读取请求体时发生错误，返回 HTTP 500 错误
		return
	}
	fmt.Println("Request Body:", string(body)) // 输出请求体内容

	// 构造响应数据并将其编码为 JSON 格式
	data, err := json.Marshal(response{
		Status:  200,
		Message: "POST request received",
		Data:    "This is a response to a POST request",
	})
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	// 设置响应头，指定内容类型为 JSON，并将编码后的数据写入响应体
	res.Header().Set("Content-Type", "application/json")
	res.Write(data)
}

func main() {
	// 注册 HTTP 处理函数，分别处理 GET 和 POST 请求
	http.HandleFunc("/get", GET)
	http.HandleFunc("/post", POST)

	// 启动 HTTP 服务器，监听在 8080 端口
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
