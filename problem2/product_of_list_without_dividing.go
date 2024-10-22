package main

import "fmt"

func product(nums []int) int {
	result := 1

	for _, i := range nums {
		result = result * i
	}
	return result
}

func main() {
	list := []int{3, 2, 4}
	l := len(list)
	leftProducts := make([]int, l)
	rightProducts := make([]int, l)

	for idx, _ := range leftProducts {
		leftProducts[idx] = product(list[:idx])
	}

	for idx, _ := range leftProducts {
		rightProducts[idx] = product(list[idx+1:])
	}

	for idx, _ := range list {
		list[idx] = leftProducts[idx] * rightProducts[idx]
	}

	fmt.Println(list)
}
