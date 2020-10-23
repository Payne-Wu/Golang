package main

import "fmt"

func main() {
	//var f = d1(10)
	fmt.Println(d1(100)(30))
}

func d1(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}
