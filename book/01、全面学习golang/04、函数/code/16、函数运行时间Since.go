package main

import (
	"fmt"
	"time"
)

func test() {
	start := time.Now()
	sum := 0
	for i := 0; i < 100000; i++ {
		sum++
	}

	elapsed := time.Since(start)
	fmt.Println("耗时", elapsed)
}

func main() {
	test()
}
