package main

import (
	"fmt"
	C "go-project/grammar/mod/calc" // 包别名
	//"go-project/grammar/mod/calc"
)

func main() {
	pkgFn()
}

func pkgFn() {
	fmt.Println(" ")
	calcAdd := C.Add(1, 3)
	fmt.Println("calcAdd", calcAdd)
}
