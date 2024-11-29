package routers

import "github.com/gin-gonic/gin"

/*
配置路由
*/

func InitRouters(r *gin.Engine) {
	BasicRoutersInit(r)
	AdminRoutersInit(r)
	ApiRoutersInit(r)
}
