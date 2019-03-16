package main

import (
	"fmt"
)

type person struct {
	name string
	id int
}

func main() {
	p := person{name: "Ramon", id: 1234}
	fmt.Println(p)
}