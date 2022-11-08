package main

import (
	"fmt"
)

func main() {
	// 3 + 4 =7
	n := 34

	result := n%10 + n/10
	fmt.Println(result)

	// 5 + 0 = 5
	n = 50
	result = n%10 + n/10
	fmt.Println(result)
}
