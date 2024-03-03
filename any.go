package main

import "fmt"

func isnemeric(s string) bool {
	str := []rune(s)
	for _, r := range str {
		if r >= '0' && r <= '9' {
			return true
		}
	}
	return false
}
func Any(f func(string) bool, a []string) bool {
	for _, char := range a {
		if f(char) {
			return true
		}
	}
	return false

}
func main() {
	a := []string{"lij", "r9", "khikh", "hh"}
	result1 := Any(isnemeric, a)
	fmt.Println(result1)
}
