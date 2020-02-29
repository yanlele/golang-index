package main

import "fmt"

type Point struct {
	x int
	y int
}

func (p Point) Add(other Point) Point {
	return Point{p.x + other.x, p.y + other.y}
}

func main() {
	p1 := Point{1, 1}
	p2 := Point{2, 2}
	result := p1.Add(p2)
	fmt.Println(result)
}

/*
结果：{3 3}
*/
