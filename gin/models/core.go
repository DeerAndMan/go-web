package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
数据库连接
https://gorm.io/zh_CN/docs/connecting_to_the_database.html
*/

var DB *gorm.DB
var DBErr error

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// 本地测试数据库
	dsn := "root:123456789@tcp(localhost:3306)/energytest?charset=utf8mb4&parseTime=True&loc=Local"
	DB, DBErr = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if DBErr != nil {
		fmt.Println("sql db err", DBErr)
	}

	fmt.Println(DB, DBErr)
}
