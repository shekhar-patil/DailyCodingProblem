package main

import "fmt"

func main() {
	list := []int{9, 15, 1, 8}
	k := 10
	flag := false

	myMap := make(map[int]int)

	for _, num := range list {
		if val, ok := myMap[num]; ok {
			if val == k-num {
				flag = true
				break
			}
		} else {
			myMap[k-num] = num
		}
	}

	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("False")
	}
}
