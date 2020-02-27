package main

import (
	"container/list"
	"fmt"
)

func main() {
	myList := list.New()

	// 最后面插入数据
	myList.PushBack("yanle")
	// 最前面插入数据
	myList.PushFront("name")

	// 最后面插入数据之后拿到句柄
	element := myList.PushBack("lele")

	// 指定元素之后插入数据
	myList.InsertAfter("insertAfter", element)

	// 指定元素之前插入数据
	myList.InsertBefore("insertBefore", element)

	// 删除指定元素
	myList.Remove(element)

	// 遍历
	for i := myList.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
