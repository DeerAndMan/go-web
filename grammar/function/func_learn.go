package main

func main() {
	var a int = 100
	var b int = 200
	var ret int
	ret = maxFunc(a, b)
	println("ret max func", ret)

	var c, d int = 1, 2
	println("before c, d", c, d)
	swap(&c, &d)
	println("c, d", c, d)
}

/*
max 测试函数
*/
func maxFunc(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

/*
!!! 不推荐
传参测试，引用传递
交换两个指针指向的值
*/
func swap(x, y *int) {
	var temp int
	temp = *x // 保存x的值
	*x = *y   // 将y的值赋给x
	*y = temp // 将temp的值赋给y
}
