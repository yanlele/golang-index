package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printString(str string, ch chan string) {
	defer wg.Done()
	fmt.Println(str)
	ch<-str
}

func main() {
	ch := make(chan string)
	wg.Add(4)
	go printString("yanle", ch)
	go printString("lele", ch)
	go printString("123", ch)
	go printString("hello", ch)

	ch<-"1"
	<-ch
	wg.Wait()
}
