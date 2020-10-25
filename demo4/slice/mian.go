package main

import "fmt"

func main() {
	//a.txt := [...]int{1, 2, 3, 4}
	//slice()
	slice1()

}

func slice() {
	// 切片的定义 -> 底层为数组， 【】中不需要定义长度
	var s1 []int
	var s2 []string
	s1 = []int{1, 2, 3}
	s2 = []string{"1", "2", "3"}
	fmt.Println(s1)
	fmt.Println(s2)
	// len长度  cap 容量
	fmt.Printf("S1  %d  %d\n", len(s1), cap(s1))
	fmt.Printf("S2  %d  %d\n", len(s2), cap(s2))

}

func slice1() {
	// 由数组得到切片 len -> 所存储到长度， cap -> 从开头切片到数组末尾
	a := [...]int{1, 2, 3, 4, 5, 6, 7}
	s1 := a[1:3]
	fmt.Printf("S1  %d  %d\n", len(s1), cap(s1))
}
