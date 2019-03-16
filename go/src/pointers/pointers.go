package main

import (
	"fmt"
)

func main() {
	i := 7
	doesnotincr(i)
	fmt.Println("i := 7, doesnotincr(i): ", i)
	incr(&i)
	fmt.Println("i := 7, incr(i): ", i)
}

func doesnotincr(x int) {
	x++
}

func incr(x *int) {
	*x++
}