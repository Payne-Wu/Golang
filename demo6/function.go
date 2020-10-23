package main

import "fmt"

/*
匿名函数
f1 := func(x, y int) {
	fmt.Println(x + y)
}


*/
func main() {
	Anonymous()
	Execute()
	Closure(1, 2)
}

func Anonymous() {
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 20)
}

func Execute() {
	// 立即执行函数
	func(x, y int) {
		fmt.Println(x + y)
	}(1, 2)
}

func Closure(x, y int) func(int, int) {
	return func(x, y int) {
		fmt.Println(x + y)
	}
}
