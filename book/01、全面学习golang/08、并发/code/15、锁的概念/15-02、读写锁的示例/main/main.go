package main

import (
	"fmt"
	"sync"
)

var (
	count int
	countGuard sync.RWMutex
)

func getCount() int {
	countGuard.RLock()
	defer countGuard.RUnlock()
	return count
}

func setCount(a int) {
	countGuard.Lock()
	defer countGuard.Unlock()
	count = a
}

func main() {
	setCount(2)
	fmt.Println(getCount())
}
