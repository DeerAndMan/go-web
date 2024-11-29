package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// MiddleW 默认中间件写法
func MiddleW(c *gin.Context) {
	fmt.Println("middleW 中间件函数")
	// 表示跳过当前中间件剩余内容，去执行下一个中间件
	//c.Next()
}

// MiddleW2
/*
中间件 第二种写法。需要执行
*/
func MiddleW2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleW2 中间函数")
		//c.Next()
	}
}
