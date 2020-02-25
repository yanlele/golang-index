package main

import "fmt"

func main() {
	var arr [30]int
	for i := 0; i < 30; i++ {
		arr[i] = i + 1
	}

	fmt.Println(arr)

	fmt.Println(arr[:20])
	fmt.Println(arr[20:])

	fmt.Println(arr[:])
	fmt.Println(arr[0:0])
}

/*
[1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30]
[1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20]
[21 22 23 24 25 26 27 28 29 30]
[1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30]
[]
*/
