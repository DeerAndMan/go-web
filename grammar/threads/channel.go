package main

import "fmt"

func channelMain() {
	basicChannel()
}

// 基础管道
func basicChannel() {
	// 创建channel，长度为5
	ch := make(chan int, 5)

	// 给管道里面存储数据
	ch <- 10
	ch <- 22
	ch <- 55

	// 获取管道里面的内容
	a := <-ch
	fmt.Println("a", a) // 10
	<-ch                // 从管道里面取值，无赋值
	b := <-ch
	fmt.Println("b", b) // 55
	ch <- 90
	ch <- 101

	// 管道长度与容量
	fmt.Printf("值：%v 容量：%v 长度：%v \n", ch, cap(ch), len(ch)) // xxx 5 2

	// 管道类型 == 引用类型
	// 管道阻塞

}

func channelFor() {
	ch1 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch1 <- i
	}
	// 使用 for range 遍历管道，当管道被关闭的时候就会退出 for range，如果没有关闭管道就会报错
	// fatal error: all goroutines are asleep - deadlock!

	//为什么不关闭通道会导致问题？
	//range 等待新数据：未关闭通道时，range 会一直阻塞等待新数据。如果数据源不再发送数据但也不关闭通道，程序就会陷入死锁。
	close(ch1)

	// channel for range 的时候是没有 key 的
	for v := range ch1 {
		fmt.Println("channel v", v)
	}
}
