package main

import (
	"fmt"
)

func main() {
	testArray := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//bestResult := []int{4, -1, 2, 1}
	//bestSum := 6

	resultArray, resultValue := getMaxSubarray(testArray)
	fmt.Println(resultArray, "\n", resultValue)
}

func getMaxSubarray(array []int) ([]int, int) {
	largestSum := -99 //min int is ideal here... placeholder to avoid fetching a const
	//bestArray := []int{}

	bestSumUsingCurrentIndex := largestSum
	bestStartingIndex := 0
	bestEndIndex := 0
	currentStartingIndex := 0
	currentEndIndex := 0

	for index, value := range array {
		if bestSumUsingCurrentIndex+value > value {
			currentEndIndex = index
			bestSumUsingCurrentIndex += value
		} else {
			currentStartingIndex = index
			bestSumUsingCurrentIndex = value
		}
		if largestSum < bestSumUsingCurrentIndex {
			largestSum = bestSumUsingCurrentIndex
			bestStartingIndex = currentStartingIndex
			bestEndIndex = currentEndIndex
		}
		largestSum = max(largestSum, bestSumUsingCurrentIndex)

	}

	return array[bestStartingIndex : bestEndIndex+1], largestSum

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
