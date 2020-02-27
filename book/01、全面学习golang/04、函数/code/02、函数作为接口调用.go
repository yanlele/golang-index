package main

import "fmt"

// 调用器接口
type Invoker interface {
	// 实现一个Call方法
	Call(interface{})
}

// 结构体类型
type Type struct {
}

// 实现 Invoker 的 Call 方法
func (s *Type) Call(p interface{}) {
	fmt.Println("form struct: ", p)
}

// 函数定义为类型
type FuncCaller func(interface{})

// 实现 Invoker 的 Call 方法
func (f FuncCaller) Call(p interface{}) {
	f(p)
}

func main() {
	// 申明接口变量
	var invoker Invoker

	// 实例化结构体
	s := new(Type)

	// 讲师李华的结构体赋值到接口
	invoker = s

	// 调用实例化结构体的方法 Type.Call
	invoker.Call("hello")

	// 匿名函数转为 FuncCaller 类型， 之后赋值给接口
	invoker = FuncCaller(func(value interface{}) {
		fmt.Println("form function: ", value)
	})

	// 使用接口调用 FuncCaller.call 内部会调用函数本体
	invoker.Call("hello")
}
