package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1) // 安全的对counter 加1
		runtime.Gosched()
	}
}

func main() {
	wg.Add(2)
	go incCounter(1)
	go incCounter(2)
	wg.Wait() // 等待goroutine结果
	fmt.Println(counter)
}
