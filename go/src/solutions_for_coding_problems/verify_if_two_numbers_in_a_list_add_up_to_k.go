package main

import (
	"fmt"
	"sort"
)

func main() {
	k := 17

	list1 := []int{10, 15, 3, 7}
	var contains_sum = containsPairWithSum(list1, k)
	fmt.Println("Contains pair with sum list={10, 15, 3, 7}, k=17? ", contains_sum)

	var list2 = []int{10, 15, 3, 6}
	contains_sum = containsPairWithSum(list2, k)
	fmt.Println("Contains pair with sum list={10, 15, 3, 6}, k=17? ", contains_sum)
}

func containsPairWithSum(list []int, k int) bool {
	sort.Ints(list)

	for i, j := 0, len(list) - 1; i < j; {
		sum := list[i] + list[j]
		if sum < k {
			i++
		} else if sum > k {
			j--
		} else {
			return true
		}
	}
	return false
}