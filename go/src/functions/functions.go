package main

import (
	"fmt"
	"errors"
	"math"
)

func main() {
	result := sum(2, 3)
	fmt.Println("sum(2, 3):", result)

	sqrt_value, err := sqrt(64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sqrt(64):", sqrt_value)
	}
}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	// It is a good practice to return errors like this because Golang does not have exceptions as other programming languages
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}