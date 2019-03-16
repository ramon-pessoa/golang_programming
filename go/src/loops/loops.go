package main

import (
	"fmt"
)

func main() {

	// The only type of loop on Golang is for

	// for loop
	fmt.Println("For loop")
	for i:= 0; i < 5; i++ {
		fmt.Println(i)
	} 

	// The for loop also works as a while loop
	fmt.Println("While loop")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// Get index and value of an array using range
	fmt.Println("Get index and value of an array using range")
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

	// Get key and value of a map using range
	fmt.Println("Get key and value of a map using range")
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"
	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}
}