package basicController

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
基础路由 控制层
*/

// BasicController 基础结构体，用于继承
type BasicController struct {
}

func (con BasicController) Index(c *gin.Context) {
	c.String(http.StatusOK, "我是来自于 controllers 控制层的结构体方法 222333")
}

func (con BasicController) Test(c *gin.Context) {
	c.String(http.StatusOK, "我是来自于 controllers 控制层的结构体方法")
}
