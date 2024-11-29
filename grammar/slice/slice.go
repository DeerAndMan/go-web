package main

import "fmt"

func main() {
	baseSlice()
	arrToSlice()
	sliceLenCap()
	sliceFun()
	makeFun()
}

func baseSlice() {
	var sliceArr = []int{1, 2}
	fmt.Printf("切片的基本数据 %v 类型：%T 长度: %v \n", sliceArr, sliceArr, len(sliceArr))

	// 扩容方法
	sliceArr = append(sliceArr, 21)
	fmt.Printf("切片扩容 %v 类型：%T 长度: %v \n", sliceArr, sliceArr, len(sliceArr))
}

func arrToSlice() {
	fmt.Println(" ")
	var arr = [5]int{55, 56, 57, 58, 59}
	fmt.Printf("arrToSlice -->> 值是什么：%v 类型：%T \n", arr, arr) // 数组
	b := arr[:]
	fmt.Printf("值是什么：%v 类型：%T \n", b, b) // 切片

	// 获取切片中的数
	c := b[1:4] // 左包，右不包
	fmt.Println("c: ", c)
}

/*
切片的长度与容量
*/
func sliceLenCap() {
	fmt.Println(" ")
	s := []int{1, 2, 3, 4, 5, 6} // 6 6
	fmt.Printf("切片的长度:%v, 容量: %d \n", len(s), cap(s))

	a := s[2:] // 4 4
	fmt.Printf("a 的长度:%v, 容量: %d \n", len(a), cap(a))

	b := s[1:3] // 2 5
	fmt.Printf("b 的长度:%v, 容量: %d \n", len(b), cap(b))

	f := s[:3] // 3 6
	fmt.Printf("f 的长度:%v, 容量: %d \n", len(f), cap(f))
}

/*
切片合并
*/
func sliceFun() {
	fmt.Println(" ")
	sliceA := []string{"php", "golang"}
	sliceB := []string{"nodejs"}
	fmt.Printf("sliceFun sliceA:%v 长度:%v, 容量: %d \n", sliceA, len(sliceA), cap(sliceA))
	fmt.Printf("sliceFun sliceB:%v 长度:%v, 容量: %d \n", sliceB, len(sliceB), cap(sliceB))

	// append
	sliceA = append(sliceA, "c++")
	fmt.Printf("添加后 sliceA:%v 长度 :%v, 容量: %d \n", sliceA, len(sliceA), cap(sliceA))

	sliceA = append(sliceA, sliceB...)
	fmt.Printf("结构 sliceA:%v 长度 :%v, 容量: %d \n", sliceA, len(sliceA), cap(sliceA))

	// copy 方法
	copyA := []int{1, 2, 3}
	copyB := make([]int, len(copyA))
	copy(copyB, copyA)
	fmt.Printf("copy方法 copyB:%v 长度 :%v, 容量: %d \n", copyB, len(copyB), cap(copyB))
}

/*
make函数
*/
func makeFun() {
	fmt.Println(" ")
	var a = "你好golang"
	s1 := []rune(a)
	fmt.Printf("rune方法 s1:%d 长度 :%v, 容量: %d \n", s1, len(s1), cap(s1))
}
