package main

import (
	"fmt"
	"strconv"
)

type Student2 struct {
	Name string
	Age  int
}

func newStudent2(age int, name string) Student2 {
	return Student2{
		Name: name,
		Age:  age,
	}
}

func (s Student2) getName2() string {
	return s.Name
}

func (s Student2) getAge2() int {
	return s.Age
}

func main() {
	ns := newStudent2(27, "yanle")
	fmt.Println(ns.getName2())
	fmt.Println(strconv.Itoa(ns.getAge2()))
}
