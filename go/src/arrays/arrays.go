package main

import (
	"fmt"
)

func main() {
	var a [5]int
	a[2] = 7

	fmt.Println(a)

	// Other way to use array
	b := [5]int{5, 4, 3, 2, 1}

	fmt.Println(b)
}