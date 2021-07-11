package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 4)
	fmt.Println(len(ch))

	ch <- 1
	ch <- 3
	ch <- 5
	ch <- 7
	fmt.Println(len(ch))

	var arr []int
	for i := 0; i < len(ch); i++ {
		fmt.Println("run")
		fmt.Println("ch: ", <-ch)
		//arr = append(arr, <-ch)
	}

	for index, value := range arr {
		fmt.Println("index: ", index)
		fmt.Println("value: ", value)
	}
}

/*
结果
0
3
*/
