package routers

import (
	"go-project/gin/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Json struct {
	Code uint8       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ListItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type JsonData struct {
	Name  string     `json:"name"`
	List  []ListItem `json:"list"`
	Total int        `json:"total"`
}

/*
后台路由组
*/

func AdminRoutersInit(r *gin.Engine) {
	AdminRouters := r.Group("admin")
	// 请求 中间件
	AdminRouters.Use(middleware.JsonResponseMiddleware)
	{
		// 根路由
		AdminRouters.GET("/", func(c *gin.Context) {
			c.String(200, "值：%v", "你好后台路由")
		})

		AdminRouters.GET("/json", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"msg":     "你好我是JSON",
			})
		})

		AdminRouters.GET("/map", func(c *gin.Context) {
			jsonData := JsonData{
				Name: "Sample Data",
				List: []ListItem{
					{ID: "1", Name: "Item 1"},
					{ID: "2", Name: "Item 2"},
				},
				Total: 2,
			}
			// 设置业务数据
			c.Set("response", jsonData)
			// 仅设置状态码，中间件会自动处理响应
			//c.Status(http.StatusOK)
			// 返回响应数据
			//c.JSON(http.StatusOK, jsonData)

			//response := Json{
			//	Code: 200,
			//	Msg:  "Success",
			//	Data: jsonData,
			//}

			//// 设置响应数据，以便中间件能够获取到它
			//c.Set("response", jsonData)
			//c.JSON(http.StatusOK, jsonData)
		})

	}
}
