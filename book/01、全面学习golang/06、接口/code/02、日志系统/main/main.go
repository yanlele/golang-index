package main

import (
	"fmt"
	. "go-index/book/01、全面学习golang/06、接口/code/02、日志系统/core"
)

// 创建日志器
func createLogger() *Logger {
	l := NewLogger()

	cw := NewConsoleWriter()

	l.RegisterWriter(cw)

	fw := NewFileWriter()

	if err := fw.SetFile("log.log"); err != nil {
		fmt.Println(err)
	}

	l.RegisterWriter(fw)
	return l
}

func main() {
	l := createLogger()
	l.Log("hello")
}
