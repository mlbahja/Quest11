package main

import "fmt"

func f(a, b int) int {
	if a == b {
		return 0
	}
	if a > b {
		return 1
	}
	return -1

}
func issorted(f func(a, b int) int, a []int) bool {
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			return false
		}
	}
	return true
}
func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	result := issorted(f, a)
	fmt.Println(result)
}
