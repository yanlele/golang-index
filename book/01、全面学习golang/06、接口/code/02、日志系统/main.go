package main

import "fmt"

// 创建日志器
func createLogger() *Logger {
	l := newLogger()

	cw := newConsoleWriter()

	l.RegisterWriter(cw)

	fw := newFileWriter()

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
