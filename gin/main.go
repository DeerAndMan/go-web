package main

import (
	"fmt"
	"go-project/gin/routers"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func main() {
	color.Red("gin 框架 - http://localhost:9981")

	// 创建一个默认的路由引擎
	r := gin.Default()

	//// 中间件
	// r.Use(middleware.MiddleW, middleware.MiddleW2())

	// 分组路由
	routers.InitRouters(r)

	// 启动一个web服务
	runErr := r.Run(":9981")

	if runErr != nil {
		fmt.Println("gin 启动失败！！！")
	}
}

func TestRouter(r *gin.Engine) {
	//// 配置路由
	//r.GET("/", func(c *gin.Context) {
	//	c.String(200, "值：%v", "你好golang gin")
	//})
	//
	//r.GET("/ping", func(c *gin.Context) {
	//	// http.StatusOK 200 成功状态码
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong -> air-verse/air 启动",
	//	})
	//})
}
