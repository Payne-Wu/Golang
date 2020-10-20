package main

import "fmt"

// if ... else if ... else
// switch
func main() {

	forLoop()
	//switchLoop()
	//gotoLoop() -> 跳至标签， break -> 退出当前循环
}

func gotoLoop() {
	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			if j == 2 {
				goto breakTag
			}
			fmt.Println(i)
		}
	}
	return
breakTag:
	fmt.Println("退出")

}

func switchLoop() {
	n := 1
	switch n {
	case 1:

	}
}

func forLoop() {
	//a := 1
	//if a == 1 {
	//	fmt.Println("GOOD")
	//} else if a == 2 {
	//	fmt.Println("Bad")
	//} else {
	//	fmt.Println("Oh~")
	//}
	if a := 1; a == 2 {
		fmt.Println(a)
	} else if a == 3 {
		fmt.Println("None")
	} else {
		fmt.Println("*")
	}
}
