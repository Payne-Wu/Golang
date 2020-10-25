package main

import "fmt"

// 自定义类型 和类型别名
type Person struct {
	name, sex, address string
	year int

}



func main() {
	p1 := Person{
		name: "Payne",
		sex: "male",
		address: "cs",
		year: 21,
	}
	p1.name = "Tom"
	fmt.Println(p1)
	
}
