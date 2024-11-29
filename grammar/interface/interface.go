package main

import "fmt"

type Animals interface {
	GetName() string
	SetName(string)
}

func main() {
	fmt.Println("interface 接口 --->>>")

	var dog1 Animals = &Dog{
		Name: "小黑",
	}

	c := &Cat{
		Name: "小美",
	}
	var cat1 Animals = c
	fmt.Println("Dog 类型数据", dog1)
	cat1.GetName()
	c.Say()
	c.SetName("不是")
	fmt.Println("cat1, c", cat1, c)

	emptyFn()
}

// Dog
/* 定义一个狗的结构体 */
type Dog struct {
	Name string
}

func (d *Dog) GetName() string {
	fmt.Println("get dog name")
	return d.Name
}
func (d *Dog) SetName(name string) {
	d.Name = name
}

// Cat
/* 猫猫结构体 */
type Cat struct {
	Name string
}

func (c *Cat) GetName() string {
	fmt.Println("get cat name")
	return c.Name
}
func (c *Cat) SetName(name string) {
	c.Name = name
}

func (c *Cat) Say() {
	fmt.Println("i`m li mei yi")
}

/*
接口细节
*/
func emptyFn() {
	fmt.Println(" ")
	var userInfo = make(map[string]interface{})
	userInfo["name"] = "张三"
	userInfo["age"] = 18
	userInfo["phone"] = "111123678"
	userInfo["hobby"] = []string{"吃饭", "睡觉", "写代码"}

	// 空接口无法获取切片，结构体中的数据
	//userInfo["hobby"][1] 错误写法
	hobby1, sucH := userInfo["hobby"].([]string)
	fmt.Println(sucH, "sucH")
	if sucH {
		fmt.Println("hobby1 获取值", hobby1[0])
	}
}
