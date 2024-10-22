package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type List struct {
	head *Node
}

func (l *List) insert(data int) {
	node := &Node{Data: data}

	if l.head == nil {
		l.head = node
		return
	}

	ptr := l.head

	for ptr.Next != nil {
		ptr = ptr.Next
	}

	ptr.Next = node
}

func (l *List) maxLength() int {
	ptr := l.head
	len := 0
	for ptr != nil {
		len++
		ptr = ptr.Next
	}

	return len
}

func findIntersection(l1 List, l2 List) *Node {
	len1 := l1.maxLength()
	len2 := l2.maxLength()

	var ptr1 *Node
	var ptr2 *Node
	var extras int

	if len1 > len2 {
		extras = len1 - len2

		ptr1 = l1.head
		i := 0
		for i < extras {
			ptr1 = ptr1.Next
			i++
		}

		ptr2 = l2.head
	} else {
		extras = len2 - len1

		ptr1 = l2.head
		i := 0
		for i < extras {
			ptr1 = ptr1.Next
			i++
		}

		ptr2 = l1.head
	}

	for ptr1.Data != ptr2.Data {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	return ptr1

}

func main() {
	s1 := []int{2, 3, 4, 8, 9, 0}
	s2 := []int{1,1,1,1,1,13, 1, 9, 0}

	l1 := List{}
	l2 := List{}

	for _, num := range s1 {
		l1.insert(num)
	}

	for _, num := range s2 {
		l2.insert(num)
	}

	node := findIntersection(l1, l2)

	fmt.Println(node.Data)

}
