package main

import (
	"fmt"
	"reflect"
)

//var 关键字 -> 声明变量（一般较多使用与全局声明） var name type
//声明
//var a string
// 批量声明 var a, b, c int
//var (
//	//b string
//	//c int
//	d []int
//	//e float32
//)
// 类型推导声明
//var a = "1"
//const name = 1
// const 声明常量(不可被更改) 与var相同
// iota： 当const 关键字出现的时候 初始化为0
func main() {
	fmt.Println(reflect.TypeOf("1"))
}
