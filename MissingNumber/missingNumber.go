package main

import "fmt"

func main() {
	testArray := []int{0, 1, 3}
	expectedResult := 2

	result := missingNumber(testArray)
	fmt.Println(result == expectedResult)
}

func missingNumber(array []int) int {
	arraySum := getSumOfArray(array)

	sizeOfArray := len(array)
	return (sizeOfArray * (sizeOfArray + 1) / 2) - arraySum
}

func getSumOfArray(array []int) int {
	result := 0
	for _, value := range array {
		result += value
	}
	return result
}
