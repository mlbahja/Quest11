package main

import "fmt"

func isprime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true

}
func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, 0, len(a))
	//result := make([]bool, len(a))
	for _, char := range a {
		result = append(result, f(char))
		//continue

	}
	return result
}

func main() {
	fmt.Println(isprime(4))
	a := []int{1, 2, 3, 4, 5, 6}
	b := Map(isprime, a)
	fmt.Println(b)

}
