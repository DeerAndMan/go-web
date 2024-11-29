package routers

import (
	"fmt"
	"go-project/gin/controllers/basicController"
	"go-project/gin/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
基础路由组
*/

func BasicRoutersInit(r *gin.Engine) {
	basicCtr := basicController.BasicController{}
	fmt.Println("basicCon", basicCtr)

	defaultRouters := r.Group("/")
	// 注册使用 中间件
	// defaultRouters := r.Group("/").use(middleware.MiddleW2())
	defaultRouters.Use(middleware.MiddleW)
	{
		// 配置路由，单独路由中间件
		defaultRouters.GET("/", middleware.HomeMiddle, func(c *gin.Context) {
			c.String(200, "值：%v", "你好golang gin vscode 配置")
		})

		defaultRouters.GET("/ping", func(c *gin.Context) {
			// http.StatusOK 200 成功状态码
			c.JSON(http.StatusOK, gin.H{
				"message": "pong -> air-verse/air 启动",
			})
		})

		defaultRouters.GET("/struct", basicCtr.Index)
	}
}
