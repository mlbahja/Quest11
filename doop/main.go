package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 4 {
		return
	}
	a, err1 := strconv.ParseInt(os.Args[1], 10, 64)
	b, err2 := strconv.ParseInt(os.Args[3], 10, 64)
	if err1 != nil || err2 != nil {
		return
	}
	switch os.Args[2] {
	case "+":
		result := a + b
		fmt.Println(result)

	case "-":
		result := a - b
		fmt.Println(result)
	case "/":
		if b == 0 {
			fmt.Println("NO division by 0")
			return
		}
		result := a / b
		fmt.Println(result)
	case "*":
		result := a * b
		fmt.Println(result)
	case "%":
		if b == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		result := a % b
		fmt.Println(result)

	default:
		return
	}

}
