package models

/*
操作的数据库表 jy_data
*/

type User struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Age     uint16 `json:"age"`
	Email   string `json:"email"`
	AddTime string `json:"add_time"`
}

// TableName
/*
表示配置操作数据库的表名称

默认情况表名是结构体名称的复数形式。如果我们的结构体名称定义成 User，表示这个模型默认操作的是 Users 表
*/
func (User) TableName() string {
	return "user"
}
