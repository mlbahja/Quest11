/*>>> return numbre of elements in slices of a string slice .
for which the f function return true*/

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
func countif(f func(string) bool, tab []string) int {
	count := 0
	for _, char := range tab {
		if f(char) {
			count++
		}

	}
	return count
}

func main() {

	table := []string{"hello", "2", "6"}
	result := countif(isnemeric, table)
	fmt.Println(result)

}
