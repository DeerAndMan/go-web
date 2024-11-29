package main

import "fmt"

func main() {
	basic()
	mapCycle()
}

func basic() {
	fmt.Println(" ")
	var myMap = make(map[string]string)
	// 第一种初始化
	myMap["name"] = "测试名称"
	fmt.Printf("map 引用类型数据 值：%v, %T \n", myMap, myMap)

	// 第二种初始化
	var userInfo = map[string]string{
		"name": "测试名称",
		"age":  "20",
	}
	fmt.Printf("map userInfo 引用类型数据 值：%v, 类型：%T \n", userInfo, userInfo)

	// make 创建，使用 interface表示任何类型
	var makeMap = make(map[string]interface{})
	makeMap["name"] = "混合类型"
	makeMap["age"] = 18
}

/*
map 数据循环遍历
*/
func mapCycle() {
	fmt.Println(" ")
	userInfo := map[string]string{
		"name": "测试名称",
		"age":  "20",
		"sex":  "男",
		"long": "18",
	}
	for k, v := range userInfo {
		fmt.Println("k:", k, "v:", v)
	}

	// 查找对应的数据是否存在
	val, vOk := userInfo["name"]
	fmt.Println("v:", val, "是否存在：", vOk)

	// 删除数据
	delete(userInfo, "long")

	fmt.Println("userInfo: ", userInfo)
}
