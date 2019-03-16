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

	// slices
	c := []int{1, 2, 3, 4, 5}
	c = append(c, 13)

	fmt.Println(c)

	// dictionary
	vertices := make(map[string]int)

	vertices["vertices"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	fmt.Println(vertices)

	delete(vertices, "square")

	fmt.Println(vertices)
}