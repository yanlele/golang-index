package main

import (
	"fmt"
	"time"
)

func running() {
	var times int
	for {
		times++
		fmt.Println("tick", times)
		time.Sleep(time.Second)
	}
}

// 代码执行后，命令行会不断地输出 tick，同时可以使用 fmt.Scanln() 接受用户输入。两个环节可以同时进行。
func main() {
	go running()

	var input string
	_, _ = fmt.Scanln(&input)
}
