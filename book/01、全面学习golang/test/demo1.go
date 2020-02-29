package main

import "fmt"

func main() {
	var a *int

	b := 12

	a = &b

	fmt.Println(*a)
}
