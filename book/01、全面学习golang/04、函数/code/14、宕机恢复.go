package main

import (
	"fmt"
	"runtime"
)

// 崩溃的时候需要传递上下文信息
type panicContext struct {
	function string
}

// 保护方式允许一个函数
func ProtectRun(entry func()) {
	defer func() {
		err := recover()

		switch err.(type) {
		case runtime.Error:
			fmt.Println("runtime error: ", err)
		default:
			fmt.Println("error: ", err)
		}
	}()

	entry()
}

func main() {
	fmt.Println("运行前")

	ProtectRun(func() {
		fmt.Println("手动宕机前")

		panic(&panicContext{"手动触发panic"})

		fmt.Println("手动宕机之后")
	})

	// 故意宕机
	ProtectRun(func() {
		fmt.Println("赋值宕机前")

		var a *int
		*a = 1

		fmt.Println("赋值宕机后")
	})

	fmt.Println("运行后")
}
