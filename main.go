// 声明 main 包，表明当前是一个可执行程序。
// package main 表示一个可独立执行的程序
package main

// 导入内置 fmt 包
import (
	"errors"
	"fmt"
	"unsafe"
)

var i int
var m = 100

//type myType int8

// main函数，是程序执行的入口
func main() {
	fmt.Println("Go Hello !") // 在终端打印 Go Hello !
	fmt.Println("i 数据的值：", i)
	fmt.Printf("m 的值=%d 类型=%T\n", m, m)
	println("在内存里面的存储空间", unsafe.Sizeof(m))
	err := sortFunc(5)
	if err != nil {
		fmt.Println("Error:", err)
	}

	//	bool 不允许强制转换
	var f = 1
	//var f1 = bool(f) // 无法强制转换
	fmt.Printf("f的强制转换 %v\n", f)

	var a = 'a'
	fmt.Printf("值：%v 类型：%T\n", a, a)

	forFun() // 循环测试
}

// 短变量声明
func sortFunc(f int) error {
	n := 10
	m := 200
	println(n, m)
	if f < 10 {
		return errors.New("text error")
	}
	return nil
}

// 循环测试
func forFun() {
	// 切片循环
	var arr = []string{"PHP", "JS", "GO"}
	for i, val := range arr {
		println(i, val)
	}

	// 混合字符串循环
	var str = "你好 test"
	for _, s := range str {
		println("str循环--->", string(s))
		//fmt.Printf("str循环的值=%q\n", s)
	}
}
