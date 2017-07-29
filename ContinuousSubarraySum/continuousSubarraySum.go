package main

import "fmt"

func main() {
	testInput := []int{23, 2, 4, 6, 7}
	testGoal := 6

	result, exists := continuousSubarraySum(testInput, testGoal)
	fmt.Println(exists)
	if exists {
		fmt.Println(result)
	}

}

func continuousSubarraySum(array []int, goal int) ([]int, bool) {
	if len(array) < 2 {
		return []int{}, false
	}

	lowPointer := 0
	highPointer := 1
	currentSubarraySum := array[lowPointer] + array[highPointer]
	for highPointer < len(array) {
		if currentSubarraySum == goal {
			fmt.Println(lowPointer, " ", highPointer)
			return array[lowPointer : highPointer+1], true
		} else if currentSubarraySum < goal {
			highPointer++
			if highPointer < len(array) {
				currentSubarraySum += array[highPointer]
			}
		} else if currentSubarraySum > goal {
			currentSubarraySum -= array[lowPointer]
			lowPointer++
			if lowPointer >= highPointer {
				highPointer++
				if highPointer < len(array) {
					currentSubarraySum += array[highPointer]
				}
			}
		}
	}
	return array[lowPointer : highPointer+1], false
}
