package main

import (
	"fmt"
	"time"
)

func main() {
	// 打点器， 每隔500毫秒触发一次
	ticker := time.NewTicker(time.Millisecond * 500)

	// 计时器, 2秒之后触发
	stopper := time.NewTimer(time.Second * 2)

	var i int

	// 不断地检查通道状况
	for {
		select {
		case <-stopper.C:
			// 计时器到了
			fmt.Println("stop")
			goto StopHere
		case <-ticker.C:
			i++
			fmt.Println("tick: ", i)
		}
	}

StopHere:
	fmt.Println("done")
}
