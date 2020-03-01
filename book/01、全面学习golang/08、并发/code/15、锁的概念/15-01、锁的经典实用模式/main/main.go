package main

import (
	"fmt"
	"sync"
)

var (
	count int
	countGuard sync.Mutex
)

func GetCount() int {
	// 锁定
	countGuard.Lock()

	defer countGuard.Unlock()
	return count
}

func SetCount(c int) {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}

func main() {
	// 可以进行安全的并发设置
	SetCount(1)
	SetCount(3)

	fmt.Println(GetCount())
}
