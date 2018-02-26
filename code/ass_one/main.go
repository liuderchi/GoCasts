package main

import (
	"fmt"
)

func main() {
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, x := range ints {
		if x%2 == 0 {
			fmt.Printf("%v is even\n", x)
		} else {
			fmt.Printf("%v is odd\n", x)
		}
	}

	fmt.Println("\n\n", newRange(5))
}

func newRange(n int) []int {
	res := []int{}
	for x := 0; x < n; x++ {
		res = append(res, x) // NOTE append is pass by value
	}
	return res
}
