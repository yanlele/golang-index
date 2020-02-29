package main

import "fmt"

type Property struct {
	value int // 属性
}

func (p *Property) SetValue(v int) {
	p.value = v
}

func (p *Property) Value() int {
	return p.value
}

func main() {
	// 实例化
	p:= new(Property)
	p.SetValue(1000)
	fmt.Println(p.Value())
}
