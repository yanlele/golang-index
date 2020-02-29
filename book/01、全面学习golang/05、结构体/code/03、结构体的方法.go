package main

import "fmt"

type Bag struct {
	items []int
}

/*
面向过程的实现方式
func insert(b *Bag, itemid int) {
	b.items = append(b.items, itemid)
}

func main() {
	bag := &Bag{}
	insert(bag, 1001)
}
*/

func (b *Bag) insert(itemid int) {
	b.items = append(b.items, itemid)
}

func main() {
	b := new(Bag)
	b.insert(1001)
	for _, value := range b.items {
		fmt.Println(value)
	}
}
