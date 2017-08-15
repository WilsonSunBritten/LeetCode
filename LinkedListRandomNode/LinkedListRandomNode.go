package main

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
}

func main() {
	// Init a singly linked list [1,2,3]
	head := ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	solution := getRandom(head)
	fmt.Println(solution)
}

func getRandom(head ListNode) int {
	return 1
}
