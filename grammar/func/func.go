package main

import (
	"errors"
	"fmt"
)

func main() {
	sum := sunFun(1, 2)
	println("sum: ", sum)
	changeFunc(1, 2, 3, 5)

	sumFn1(1)
	// 返回多个值
	sum1, sub1 := calc(3, 5)
	fmt.Println("sum1: ", sum1, "sub1: ", sub1)

	sum2, sub2 := calc1(30, 50)
	fmt.Println("sum2: ", sum2, "sub2: ", sub2)

	// 类型定义
	typeFn()

	// 函数方法
	deferFn()

	// 错误捕捉
	recoverFn()
}

func sunFun(x int, y int) int {
	sum := x + y
	return sum
}

/*
函数的可变参数
表示该函数参数不固定
*/
func changeFunc(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println("changeFunc ===>", sum)
	return sum
}

/*
固定参数与可变参数同时出现
可变参数注意位置（在后面）
*/
func sumFn1(x int, y ...int) int {
	sum := x
	for _, v := range y {
		sum += v
	}
	return sum
}

/*
一次返回多个返回值
*/
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

/*
返回值命名：函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回
*/
func calc1(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	// 直接返回
	return
}

/*
函数类型可以使用关键字 type 来定义
*/
// 函数类型
type calculation func(int, int) int

func typeFn() {
	fmt.Println(" ")
	var c calculation
	c = add // 包里的类型
	fmt.Printf("函数类型 C：%T \n", c)
	d := add // 里面的类型是推导出来的类型
	fmt.Printf("函数类型 D：%T \n", d)

	var bFn = adder()
	fmt.Printf("闭包函数 bibaoFn： %T \n", bFn)
	fmt.Println("bFn()", bFn()) // 11
	fmt.Println("bFn()", bFn()) // 11
	fmt.Println("bFn()", bFn()) // 11

	fmt.Println(" ")
	var bFn2 = adder2()
	fmt.Println("bFn2()", bFn2(1)) // 11
	fmt.Println("bFn2()", bFn2(1)) // 12
	fmt.Println("bFn2()", bFn2(1)) // 13

}
func add(x, y int) (sum int) {
	sum = x + y
	return
}

/*
闭包：有权访问另一个函数作用域中的变量的函数
*/
func adder() func() int {
	var i = 10
	return func() int {
		return i + 1
	}
}
func adder2() func(y int) int {
	// 内存常驻，无法 GC
	var i = 10
	return func(y int) int {
		i += y
		return i
	}
}

/*
defer 延迟执行
*/
func deferFn() {
	fmt.Println(" ")
	fmt.Println("f1 defer 返回的值：", f1()) // 0
	fmt.Println("f2 返回的值：", f2())       // 1

	type Test struct {
		name string
	}
	ts := []Test{{"a"}, {"b"}, {"c"}}
	fmt.Println("f2 返回的值：", ts)
}
func f1() int {
	var a int = 0
	defer func() {
		a++
	}()
	return a
}

func f2() (a int) {
	defer func() {
		a++
	}()
	return a
}

/*
recover
*/
func recoverFn() {
	fmt.Println(" ")
	fnC3(10, 0)
	fnC3(10, 2)

	err := readFile("test")
	fmt.Println("err panic:", err)
	if err != nil {
		panic(err)
	}
}
func fnC3(a, b int) {
	// 异常处理
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("recover 错误：", err)
		}
	}()
	fmt.Println(a / b)
}

/*
读取文件
*/
func readFile(fileName string) error {
	if fileName == "test.go" {
		return nil
	} else {
		return errors.New("文件读取失败！！！")
	}
}
