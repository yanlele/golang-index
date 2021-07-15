## 锁的研究


### 锁的作用？
- 互斥锁 
- 读写锁


#### 举一个简单的例子

```go
package main

import (
	"fmt"
	"sync"
)

// 全局变量
var num int64
var wg sync.WaitGroup

func add() {
	for i := 0; i < 10000000; i++ {
		num = num + 1
	}
	// 协程退出， 记录 -1
	wg.Done()
}
func main() {
	// 启动2个协程，记录 2
	wg.Add(2)

	go add()
	go add()

	// 等待子协程退出
	wg.Wait()
	fmt.Println(num)
}
```

按照上述代码，我们的输出结果应该是 20000000，每一个协程计算 10000000 次，
可是实际结果却是 `10334141`

每一次计算的结果还不一样，出现这个问题的原因就是上述提到的资源竞争

两个 goroutine 协程在访问和修改num变量，会存在2个协程同时对num+1 ， 最终num 总共只加了 1 ，而不是 2


### 互斥锁
每个对象都对应于一个可称为互斥锁的标记，这个标记用来保证在任一时刻，只能有一个协程访问该对象。

场景： 写大于读操作的

#### 解决问题
- sync包 的 Mutex类型 来实现互斥锁
```go
package main

import (
	"fmt"
	"sync"
)

/*互斥锁*/

var num int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 10000000; i++ {
		// 访问前给资源加锁
		lock.Lock()
		num = num + 1
		lock.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)

	go add()
	go add()

	wg.Wait()
	fmt.Println(num)
}
```

### 读写锁
若我们并发的去读取一个资源，且不对资源做任何修改的时候如果也要加锁才能读取数据，是不是就很没有必要呢                  
这种场景下读写锁就发挥作用了，他就相对灵活了，也很好的解决了读多写少的场景问题                     

简单点儿说就是： 读的时候，不能写；写的时候， 不能读


读写锁的种类：                     
- 读锁
- 写锁

demo:               
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

/*读写锁*/

var (
	num    int64
	wg     sync.WaitGroup
	rwLock sync.RWMutex
)

func write() {
	// 加上写锁
	rwLock.Lock()

	num = num + 1

	time.Sleep(10 * time.Millisecond)

	rwLock.Unlock()

	wg.Done()
}

func read() {
	rwLock.RLock()

	time.Sleep(time.Millisecond)

	rwLock.RUnlock()

	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end:=time.Now()
	fmt.Println(end.Sub(start))
}
```

简单来说，在并发过程中，若其中一个协程拿不到锁，他会不停的去尝试拿锁，
不停的去看能不能拿，而不是阻塞睡眠


### 如何选择锁？
若写的频次大大的多余读的频次，那么选择互斥锁
若读的频次大大的多余写的频次，那么选择读写锁


### 原子操作
原子操作是不可分割的，在执行完毕之前不会被任何其它任务或事件中断

上述我们的加锁案例，咱们编码中的加锁操作会涉及内核态的上下文切换会比较耗时、代价比较高                 
针对基本的数据类型我们还可以使用原子操作来保证并发安全                     
因为原子操作是Go语言提供的方法它在用户态就可以完成，因此性能比加锁操作更好                  
不用我们自己写汇编，这里 GO 也提供了原子操作的包，供我们一起来使用 sync/atomic                     

```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/* 原子操作的示例 */

var num int64
var lock sync.Mutex
var wg sync.WaitGroup

// 普通版本
func add() {
	num = num + 1
	wg.Done()
}

// 互斥锁版本
func mutexAdd() {
	lock.Lock()
	num = num + 1
	lock.Unlock()
	wg.Done()
}

func atomicAdd() {
	atomic.AddInt64(&num, 1)
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 20000; i++ {
		wg.Add(1)
		//go add()
		//go mutexAdd()
		go atomicAdd()
	}
	wg.Wait()

	end := time.Now()
	fmt.Println("num: ", num)
	fmt.Println(end.Sub(start))
}
```




### 参考文档
- [GO的锁和原子操作分享](https://juejin.cn/post/6972846349968474142)
- [代码参考](https://github.com/yanlele/go-index-core/tree/master/demos/21%E5%B9%B4/07%E6%9C%88/02%E3%80%81%E9%94%81%E7%9A%84%E7%A0%94%E7%A9%B6)


