package main

import "fmt"

func main() {
	//二者都是用来做内存分配的。
	//make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
	//而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。

	//base()

	New()

}

func base() {
	// &取地址值 *根据地址取值
	n := 18
	fmt.Println(&n)
	p := &n
	fmt.Printf("%d\n", *p)

}

func New() {
	// 错误写法 初始值为nil，无法根据内存地址取值
	//var a *int只是声明了一个指针变量a但是没有初始化，指针作为引用类型需要初始化后才会拥有内存空间，才可以给它赋值。
	/*
		var a *int
			//fmt.Println(a) // nil
			*a = 100
			fmt.Println(*a) // panic: runtime error: invalid memory address or nil pointer dereference
	*/

	// new 申请一个内存地址值
	var a2 = new(int)
	*a2 = 100
	fmt.Println(a2)     //0xc00001a080
	fmt.Println(*a2)    //100
	fmt.Println(&(*a2)) // 0xc00001a080
	*a2 = 300
	fmt.Println(a2)  //0xc0000b2008
	fmt.Println(*a2) // 300

}

func Make() {
	/*
			make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，
			而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，
			所以就没有必要返回他们的指针了
		make(t Type, size ...IntegerType) Type
	*/

	var a map[string]int
	a = make(map[string]int, 10)
	a["NaZha"] = 100
	fmt.Println(a)

}
