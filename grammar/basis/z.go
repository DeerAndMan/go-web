package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
)

/* & 和 * 运算符实例 */
func main() {
	var a int = 4
	var ptr *int // 定义整型指针变量

	ptr = &a // 把变量a的地址赋给ptr整型指针变量
	fmt.Printf("*ptr的值 为 %d \n", *ptr)

	var b int = rand.Intn(100)
	if b > 10 {
		fmt.Println("a 大于 b", b)
	} else {
		println("a 小于 b", b)
	}

	var c int = rand.Intn(22)
	switch {
	case c >= 0 && c <= 10:
		println("c 在 0 - 10 之间")
	case c > 10 && c <= 15:
		println("c 在 10 - 15 之间")
	case c > 15 && c < 20:
		println("c 在 15 - 20 之间")
	default:
		//fmt.Println("c 不在任何指定范围内")
		color.Red("c 不在任何指定范围内")
	}

	// case 允许存在多个判断值
	switch score := "D"; score {
	case "a", "B", "c":
		println("及格")
	case "D":
		println("不及格")
	}

	/* for 循环 */
	sum := 0
	for i := 0; i <= 1000; i++ {
		sum += i
	}
	println("for sum", sum)

	/* for 循环 该字符对应的ASCII码值 */
	s := "abc"
	n := len(s)
	for n > 0 {
		println("递减 --->", s[n-1], "string 转化 ===", string(s[n-1]))
		n--
	}

	/* range 迭代数据结构 */
	numbers := []int{1, 2, 3, 4, 5}
	for i, num := range numbers {
		fmt.Printf("索引 %d, 值 %d \n", i, num)
	}
	myMap := map[string]int{"apple": 3, "banana": 5, "cherry": 7}
	for key, value := range myMap {
		fmt.Printf("键: %s, 值: %d\n", key, value)
	}

	/*
		break终止循环：一次跳出多重循环
	*/
label1:
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			if j == 2 {
				break label1
			}
			println("for break 循环")
		}
	}

	/* goto */
	var p int = 110
loop:
	for p < 120 {
		if p == 115 {
			p = p + 1
			goto loop
		}
		fmt.Printf("goto p 的值： %d \n", p)
		p++
	}
}
