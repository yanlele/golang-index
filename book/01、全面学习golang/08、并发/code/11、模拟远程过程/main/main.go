package main

import (
	"errors"
	"fmt"
	"time"
)

// 模拟客户端的请求和接受信息
func RPCClient(ch chan string, req string) (string, error) {
	// 向服务端发送请求
	ch <- req

	select {
	case ack := <-ch:
		// 接收到服务器的返回信息
		return ack, nil
	case <-time.After(time.Second):
		// 超时
		return "", errors.New("Time out")

	}
}

func RPCServer(ch chan string) {
	for {
		// 接受客户端的请求
		data := <-ch

		fmt.Println("server received: ", data)

		ch <- "roger"
	}
}

func main() {
	ch := make(chan string)

	go RPCServer(ch)

	recv, err := RPCClient(ch, "hi")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("client received", recv)
	}
}
