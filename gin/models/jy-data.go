package models

/*
操作的数据库表 jy_data
*/

type jyData struct {
	ID     int    `json:"id"`
	Bz     string `json:"bz"`
	Cbjg   string `json:"cbjg"`
	Cbjgex string `json:"cbjgex"`
	Ckcb   string `json:"ckcb"`
	Ckcbj  string `json:"ckcbj"`
	Ckyk   string `json:"ckyk"`
}

// TableName
/*
表示配置操作数据库的表名称
*/
func (jyData) TableName() string {

	return "test"
}
