package main

import (
	"fmt"
)

func main() {
	var x int = 5 // var x int = 5  or v := 5
	y := 7 // golang will infere the type
	var sum int = x + y

	fmt.Println(sum)
}