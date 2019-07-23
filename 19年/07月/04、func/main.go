package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	Name string
	Age  int
}

func newStudent(age int, name string) *Student {
	return &Student{
		Name: name,
		Age:  age,
	}
}

func (s *Student) getName() string {
	return s.Name
}

func (s *Student) getAge() int {
	return s.Age
}

func main() {
	ns := newStudent(27, "yanle")
	fmt.Println(ns.getName())
	fmt.Println(strconv.Itoa(ns.getAge()))
}
