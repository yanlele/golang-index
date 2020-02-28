package main

import (
	"fmt"
	"time"
)

func test() {
	start := time.Now()
	sum:= 0
	for i := 0; i < 10000000; i++ {
		sum++
	}
	elapsed := time.Now().Sub(start)
	fmt.Println(elapsed)
}

func main() {
	test()
}
