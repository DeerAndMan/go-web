package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func HomeMiddle(c *gin.Context) {
	fmt.Println("特定路由的中间件 ---> home")
	c.Next()
}
