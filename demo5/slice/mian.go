package main

import "fmt"

func main() {
	//a := [...]int{1, 2, 3, 4}
	slice()

}

func slice() {
	// 切片的定义
	var s1 []int
	var s2 []string
	s1 = []int{1, 2, 3}
	s2 = []string{"1", "2", "3"}
	fmt.Println(s1)
	fmt.Println(s2)

}
