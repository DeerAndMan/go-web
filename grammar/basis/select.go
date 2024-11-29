package main

import "fmt"

func main1() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received", i1, "from c1 \n")
	case c2 <- i2:
		println("sent", i2)
	case i3, ok := (<-c3):
		if ok {
			fmt.Printf("received", i3, "from c3 \n")
		} else {
			println("c3 is close")
		}
	default:
		println("no communicate")
	}
}
