package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3}
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for index, value := range a {
		fmt.Printf("index: %d, value: %d \n", index, value)
	}
}
