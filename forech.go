package main

import (
	"fmt"
	"math"

	"github.com/01-edu/z01"
)

func PrintNBR(nb int) {
	if nb == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		PrintNBR(223372036854775808)
		return
	}
	if nb < 0 {
		z01.PrintRune('-')
		nb = -nb
	}
	if nb >= 10 {
		PrintNBR(nb / 10)
	}
	z01.PrintRune(rune(nb%10) + '0')
}
func ForEach(f func(int), a []int) {
	for _, char := range a {
		f(char)
	}
}

func main() {
	fmt.Println(math.MaxInt)
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNBR, a)
}
