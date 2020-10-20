package main

import "fmt"

// Array 存放数据的容器 默认值为类型的默认值 eg： int -> 0, string -> ' ', bool -> false
func main() {
	//func1()
	//func2()
	func3()
}

func func1() {
	//var testArray [3]int                        //数组会初始化为int类型的零值
	//var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	//var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	var a1 [3]int
	a2 := [3]int{1, 2}
	a3 := [2]bool{true, false}
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
}

func func2() {
	//注意： 多维数组只有第一层且有具体的值才可以使用...来让编译器推导数组长度。例如：
	var a1 [3][3]int
	var a2 [4][3]int
	a3 := [3][4]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
}

func func3() {
	// 指定索引值的方式来初始化数组
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", a) //type of a:[4]int
}
