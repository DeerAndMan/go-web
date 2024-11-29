package routers

import "github.com/gin-gonic/gin"

/*
接口路由
*/

func ApiRoutersInit(r *gin.Engine) {
	ApiRouters := r.Group("api")
	{
		ApiRouters.GET("/", func(c *gin.Context) {
			c.String(200, "值：%v", "你好业务路由")
		})
	}
}
