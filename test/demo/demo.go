package main

//func main() {
//	a := []int{1, 2, 3, 4}
//	b := 1
//	println(contain(a, b))
//}
//
//func contain(a []int, b int) bool {
//	for _, x := range a {
//		if x == b {
//			return true
//		}
//	}
//
//	return false
//}
var a [10]int
func main() {
	nums := []int{2,7,11,15}
	twoSum(nums, 9)
}
func twoSum(nums []int, target int) []int {
	for x, y := range nums {
		a := target - nums[x]
		if contain(nums, a) && a != y {return []int{x, a.indexOf(a)}}
	}
	return nil
}
func contain(a []int, b int) bool {
	for _, x := range a {
		if x == b {
			return true
		}
	}

	return false
}