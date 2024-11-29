package calc

/*
自定义包
*/

var privateA = "私有变量 A"

var publicB = "公有变量 B"

func Add(x, y int) int { // 首字母大写表示公有方法
	return x + y
}

// Sub // 公有方法
func Sub(x, y int) int {
	return x - y
}
