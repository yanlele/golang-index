package main

import "fmt"

func main() {
	var arr []int
	arr = append(arr, 1)
	fmt.Println(arr)

	arr = append(arr, 2, 3, 4, 5)
	fmt.Println(arr)

	arr = append(arr, []int{6, 7, 8}...)
	fmt.Println(arr)

	arr = append([]int{0}, arr...)
	fmt.Println(arr)

	arr = append([]int{-3, -2, -1}, arr...)
	fmt.Println(arr)
}

/*
结果：
[1]
[1 2 3 4 5]
[1 2 3 4 5 6 7 8]
[0 1 2 3 4 5 6 7 8]
[-3 -2 -1 0 1 2 3 4 5 6 7 8]
*/