package main

import "fmt"

func main() {
	fmt.Println("deffer begin")

	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)

	fmt.Println("deffer end")
}
/*
执行结果
deffer begin
deffer end
3
2
1
*/