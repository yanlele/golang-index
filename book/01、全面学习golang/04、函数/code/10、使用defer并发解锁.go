package main
import "sync"

/*
如果不使用 deffer 加锁解锁需要这样处理
var (
	// 一个演示用的映射
	valueByKey = make(map[string]int)
	// 保证使用映射时的并发安全互斥锁
	valueByKeyGuard = sync.Mutex{}
)

func readValue(key string) int {
	// 对共享资源加锁
	valueByKeyGuard.Lock()

	// 取值
	v := valueByKey[key]

	// 对共享资源解锁
	valueByKeyGuard.Unlock()
	return v
}
*/

var (
	// 一个演示用的映射
	valueByKey = make(map[string]int)
	// 保证使用映射时的并发安全互斥锁
	valueByKeyGuard = sync.Mutex{}
)

func readValue(key string) int {
	// 对共享资源加锁
	valueByKeyGuard.Lock()

	// 对共享资源解锁
	defer valueByKeyGuard.Unlock()
	return valueByKey[key]
}
