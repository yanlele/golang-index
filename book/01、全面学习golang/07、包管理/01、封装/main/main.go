package main

import (
	"fmt"
	"golang-index/book/01、全面学习golang/07、包管理/01、封装/modal"
)

func main() {
	p := modal.NewPerson("smith")
	p.SetAge(18)
	p.SetSal(5000)
	fmt.Println(p)
	fmt.Println(p.Name, "age = ", p.GetAge(), " sal = ", p.GetSal())
}
