package main

import (
	"fmt"
	"reflect"
)

// 反射
func main() {
	reflection1()
}

type Person struct {
	Name string
}

// 反射获取变量的类型
func reflection1() {
	fmt.Println(" ")
	fmt.Println("反射！！！")
	a := 1
	b := "2"
	c := false
	d := Person{
		Name: "映射",
	}
	e := [3]int{1, 2, 3}
	f := []int{11, 22, 33}

	printfType(a)  // 数据类型：int 类型名称：int 类型种类：int
	printfType(b)  // 数据类型：string 类型名称：string 类型种类：string
	printfType(&c) // 数据类型：*bool 类型名称： 类型种类：ptr
	printfType(d)  // 数据类型：main.Person 类型名称：Person 类型种类：struct
	printfType(e)  // 数据类型：[3]int 类型名称： 类型种类：array
	printfType(f)  // 数据类型：[]int 类型名称： 类型种类：slice

	changeReflection(&a)
	fmt.Println("修改后A的值：", a)
}
func printfType(v interface{}) {
	x := reflect.TypeOf(v)
	fmt.Printf("数据类型：%v 类型名称：%v 类型种类：%v \n", x, x.Name(), x.Kind())
}

// 反射修改值
func changeReflection(x interface{}) {
	fmt.Println(" ")
	v := reflect.ValueOf(x) // 引用地址
	if v.Elem().Kind() == reflect.Int {
		v.Elem().SetInt(110)
	}

	fmt.Println("修改反射的值：", v, v.Kind(), v.Elem().Kind(), reflect.Int)
}
