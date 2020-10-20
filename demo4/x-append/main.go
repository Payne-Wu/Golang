package main

import "fmt"

func main() {
	appends()
}

func appends() {
	// append 为切片添加元素
	s1 := []string{"bj", "sh", "gz", "cs"}
	fmt.Printf("Len -> %d ; Cap -> %d \n", len(s1), cap(s1))
	// 调用append 必须使用原来到切片变量返回值
	s1 = append(s1, "sz")
	fmt.Println(s1)
	fmt.Printf("Len -> %d ; Cap -> %d \n", len(s1), cap(s1))
	s2 := []string{"sz", "hz"}
	s1 = append(s1, s2...)
	fmt.Println(s1)
	fmt.Printf("Len -> %d ; Cap -> %d \n", len(s1), cap(s1))

}
