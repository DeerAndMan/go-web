package addminController

import (
	"github.com/gin-gonic/gin"
	models2 "go-project/gin/models"
)

/*
控制层
*/

// UserController 结构体 继承
type UserController struct {
}

func (con UserController) Index(c *gin.Context) {
	// 查询数据库
	var userList []models2.User

	// 将查询到的数据赋值给 userList
	models2.DB.Find(&userList)

	c.JSON(200, gin.H{
		"result": userList,
	})

	c.String(200, "用户列表")
}
