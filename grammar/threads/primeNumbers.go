package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// demo
func demo() {
	fmt.Println(" ")
	start := time.Now()

	var sum []int
	for m := 2; m < 60000; m++ {
		var flag = true
		for i := 2; i < m; i++ {
			if m%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			//sum = append(sum, m)
		}
	}
	fmt.Println("demo 结束 sum --->", sum)
	fmt.Println("时长", time.Since(start))
}

// 开启协程
func startGoroutine() {
	fmt.Println(" ")
	// 获取分段范围
	ranges := calculateRanges(2, 60000, 6)
	fmt.Println("ranges", ranges)
	start := time.Now()
	var sum int = 0
	for _, r := range ranges {
		wg.Add(1)
		go findPrimes(r.start, r.end, &wg, &sum)
	}
	wg.Wait()
	fmt.Println("执行时间：", time.Since(start))
	fmt.Println("所有协程执行完毕", sum)
}

func findPrimes(start, end int, wg *sync.WaitGroup, sum *int) {
	defer wg.Done()
	fmt.Println(start, end)
	for m := start; m < end; m++ {
		if isPrime(m) {
			mutex.Lock() // 使用互斥锁安全更新 sum
			*sum++
			mutex.Unlock()
		}
	}
}

type RangeImpl struct {
	start, end int
}

// 计算数字区间，将范围分成指定的段数
func calculateRanges(startNum, endNum, segments int) []RangeImpl {
	var ranges []RangeImpl
	rangeSize := (endNum - startNum) / segments

	// 分配每个区间的范围
	for i := 0; i < segments; i++ {
		startRange := startNum + i*rangeSize + i
		// 最后一个区间需要多处理余数部分
		endRange := startRange + rangeSize
		if i == segments-1 {
			endRange = endNum // 最后一个区间处理剩余的部分
		}
		ranges = append(ranges, RangeImpl{startRange, endRange})
	}
	return ranges
}

// isPrime 判断一个数字是否是素数
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	// 只需检查到平方根
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
