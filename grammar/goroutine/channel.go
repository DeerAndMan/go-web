package main

import (
	"fmt"
)

/*
通道 channel
*/

func main() {
	var ch chan int
	// 定义一个传输整型数据的通道，默认不带缓冲区
	// 缓冲区大小为 3
	ch = make(chan int, 3)

	fmt.Printf("1. ch 的值=%v ch本身的地址=%p \n", ch, &ch)

	//ch <- v   // 把 v 发送到通道 ch
	//v := <-ch // 从 ch 接收数据，并把值赋给 v
	// 因为 ch 是带缓冲的通道，我们可以同时发送两个数据
	// 而不用立刻需要去同步读取数据
	ch <- 1
	ch <- 211
	ch <- 50
	//println("通道 1 的值：", ch)
	//println("通道 2 的值：", <-ch)
	//println("通道 3 的值：", <-ch)
	println("channel len, cap", len(ch), cap(ch))

	// 从 channel 取出数据后，可以继续放入
	<-ch
	ch <- 100
	println("通道 3 的值：", <-ch)
}
