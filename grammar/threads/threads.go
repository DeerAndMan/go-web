package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
线程 协程
*/
func main() {
	// 协程
	//coroutines()
	//getCpu()

	//demo()
	//startGoroutine()

	// 管道
	channelMain()
	channelFor()
}

// Coroutines
/*
协程 等待组
主进程执行结束，无论 协程 是否结束都会被关闭
*/
var wg sync.WaitGroup

/* 互斥锁 */
var mutex sync.Mutex

func coroutines() {
	wg.Add(1)   //协程计数器加 1
	go testFn() // 表示开启一个协程
	wg.Add(1)
	go testFn2()

	wg.Wait() // 等待 协程 执行完毕
	fmt.Println("主线程执行完毕 。。。")
}
func testFn() {
	for i := 0; i < 10; i++ {
		fmt.Println("testFn -> 你好 golang", i)
		time.Sleep(time.Millisecond * 10)
	}
	wg.Done() //协程计数器加 -1
}
func testFn2() {
	for i := 0; i < 10; i++ {
		fmt.Println("testFn2 ->  你好 golang", i)
		time.Sleep(time.Millisecond * 300)
	}
	wg.Done() //协程计数器加 -1
}

// getCpu 获取CPU的个数
func getCpu() {
	fmt.Println(" ")
	// 获取计算机上面的cpu个数
	cpuNum := runtime.NumCPU()
	fmt.Println("CPU的个数", cpuNum)

	// 可以设置使用多个 cpu
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("CPU 设置完成")
}
