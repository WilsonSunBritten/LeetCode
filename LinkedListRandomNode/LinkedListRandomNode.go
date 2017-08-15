package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ListNode struct {
	value int
	next  *ListNode
}

func main() {
	// Init a singly linked list [1,2,3]
	head := ListNode{1, &ListNode{2, &ListNode{3, nil}}}

	oneCount := 0
	twoCount := 0
	threeCount := 0

	for i := 0; i < 5000; i++ {
		solution := getRandom(head)
		switch solution {
		case 1:
			oneCount++
			break
		case 2:
			twoCount++
			break
		case 3:
			threeCount++
		}
	}
	fmt.Println("oneCount: ", oneCount, " twoCount: ", twoCount, " threeCount: ", threeCount)
}

func getRandom(head ListNode) int {
	currentSolution := head.value
	currentNode := head
	rand.Seed(time.Now().UnixNano())
	for position := 1; currentNode.next != nil; position++ {
		currentNode = *currentNode.next
		if rand.Intn(position+1) == position {
			currentSolution = currentNode.value
		}
	}
	return currentSolution
}
