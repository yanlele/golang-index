package main

import "fmt"

func Accumulate(value int) func() int {
	return func() int {
		value++
		return value
	}
}

func main() {
	accumulator := Accumulate(1)

	// 累加1并且打印
	fmt.Println(accumulator())
	fmt.Println(accumulator())

	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator)

	// 新创建一个累加器
	accumulator2 := Accumulate(10)

	// 累加
	fmt.Println(accumulator2())

	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator2)
}
/*
结果：
2
3
0xc00000e028
11
0xc00000e038
*/
