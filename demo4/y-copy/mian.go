package main

import "fmt"

// copy
func main() {
	a1 := []int{1, 3, 5, 7, 9}
	a2 := a1
	/*
		make函数：`分配`并初始化一个类型的对象
	*/
	var a3 = make([]int, 6, 6)
	/*
		copy（copy(destSlice, srcSlice []T)）  destSlice: 目标切片 srcSlice: 数据来源切片
		将后面的给前面
	*/
	copy(a3, a1)
	//copy(a1, a3)
	fmt.Println(a1, a2, a3)
}
