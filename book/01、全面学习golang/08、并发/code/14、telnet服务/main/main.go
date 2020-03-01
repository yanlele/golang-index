package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// 服务逻辑， 传入地址和退出的信道
func server(address string, exitChan chan int) {
	// 根据给定地址进行监听
	l, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Println(err.Error())
		exitChan <- 1
	}

	fmt.Println("listen: ", address)

	// 延迟关闭监听器
	defer l.Close()

	// 监听循环
	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		go handleSession(conn, exitChan)
	}
}

func handleSession(conn net.Conn, exitChan chan int) {
	fmt.Println("Session started")

	// 创建一个网络连接数据的读取器
	reader := bufio.NewReader(conn)

	// 循环接受数据
	for {
		// 多去字符串， 直到碰到回车返回
		str, err := reader.ReadString('\n')

		if err == nil {
			// 正确读取
			str = strings.TrimSpace(str)

			// 处理 Telnet 指令
			if !processTelnetCommand(str, exitChan) {
				_ = conn.Close()
				break
			}

			// Echo 逻辑， 发什么数据， 原样返回
			_, _ = conn.Write([]byte(str + "\r\n"))
		} else {
			// 发生错误
			fmt.Println("Session closed")
			_ = conn.Close()
			break
		}
	}
}

func processTelnetCommand(str string, exitChan chan int) bool {
	// @close 指令表示终止本次对话
	if strings.HasPrefix(str, "@close") {
		fmt.Println("Session closed")
		return false
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("server shutdown")
		// 往信道中写入0， 阻塞等待接收方处理
		exitChan <- 0
		return false
	}

	fmt.Println(str)

	return true
}

func main() {
	exitChan := make(chan int)

	go server("127.0.0.1:7001", exitChan)

	code := <- exitChan

	os.Exit(code)
}
