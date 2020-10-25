package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	a = append(a, 3)
	fmt.Printf("Len a.txt: %d  Cap: %d \n", len(a), cap(a))
	a = append(a[:2], a[7:]...) // del a.txt[3:7]
	fmt.Println(a)
	a = append(a, 3)
	fmt.Printf("Len a.txt: %d  Cap: %d \n", len(a), cap(a))

}
