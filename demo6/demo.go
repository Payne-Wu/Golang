package main

import (
	"fmt"
)

func main() {
	fmt.Println(fib(10))

}

func fib(n int) int {
	if n < 2 {
		return 1
	}
	if n == 2 {
		return 2
	}

	now, p1, p2 := 0, 1, 1
	for i := 3; i < n; i++ {
		now = p1 + p2
		p1 = p2
		p2 = now
		//now, p1, p2 = p1+p2, p2, now
		fmt.Printf("Now: %d, P1:%d, P2:%d \n", now, p1, p2)
	}
	return now
}
