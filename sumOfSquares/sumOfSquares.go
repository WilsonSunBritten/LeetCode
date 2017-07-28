package main

import "fmt"

func main() {
	testValue := 5

	resultAnswer, exists := sumOfSquares(testValue)
	if exists {
		fmt.Println("answer: ", resultAnswer)
	} else {
		fmt.Println("no solution")
	}
}

func sumOfSquares(c int) ([2]int, bool) {
	sqrtC := sqrt(c)

	for a := 0; a <= int(sqrtC); a++ {
		bSquared := c - a*a
		b := sqrt(bSquared)
		if isInteger(b) {
			return [2]int{a, int(b)}, true
		}
	}
	return [2]int{0, 0}, false
}

func sqrt(input int) float64 {
	floatInput := float64(input)
	if floatInput == 0 {
		return 0
	}

	guess := 1.0
	for iteration := 0; iteration < input; iteration++ {
		guess = guess - ((pow(guess, 2) - floatInput) / (2 * guess))
	}
	return guess
}

//only positive int powers
func pow(input float64, iterations int) float64 {
	result := 1.0

	for iteration := 0; iteration < iterations; iteration++ {
		result *= input
	}
	return result
}

//just tests if float is very close to integer value
func isInteger(input float64) bool {
	if input-float64(int(input)) < .0001 {
		return true
	}
	return false
}
