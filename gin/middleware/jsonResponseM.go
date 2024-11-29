package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Json struct {
	Code uint8       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

/*
通用 JSON 结构体
*/

func JsonResponseMiddleware(c *gin.Context) {
	// 在请求处理完成之后，统一格式化响应
	c.Next()

	// 如果响应已经被写入，不再处理
	if c.Writer.Written() {
		return
	}

	// 获取 c.Writer 的响应状态码和响应内容
	status := c.Writer.Status()

	// 默认返回的 Msg 和 Data
	msg := http.StatusText(status)
	var data interface{}

	// 检查是否存在错误
	if len(c.Errors) > 0 {
		status = http.StatusInternalServerError // 设置状态码为 500
		msg = c.Errors[0].Error()               // 返回第一个错误信息
	} else if response, exists := c.Get("response"); exists {
		data = response // 获取业务逻辑传递的数据
	}

	// 包装成统一格式的 JSON 响应
	c.JSON(status, Json{
		Code: uint8(status),
		Msg:  msg,
		Data: data,
	})
}
