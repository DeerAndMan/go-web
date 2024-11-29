package main

import "fmt"

func main() {
	var balance = [5]float32{100, 2.0, 3.4, 7.0, 50.0}

	fmt.Println("balance is:", balance)

	for i := 0; i < len(balance); i++ {
		balance[i] = balance[i] + 20
	}

	fmt.Println("balance for add:", balance)

	var twoDimension = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("twoDimension 二维数组：", twoDimension[1][2])

	/* 自动推导 */
	var arr2 = [...][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println("arr2 ==>", arr2)
}
