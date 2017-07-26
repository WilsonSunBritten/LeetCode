package main

import "fmt"

func main() {
	testCoins := []int{1, 2, 5}
	testAmount := 11

	result := coinChange(testCoins, testAmount)
	fmt.Println("testCoins: ", testCoins, "\ntestAmount: ", testAmount, "\nresult", result)
}

//result is an array of coin values that most efficiently add up to  amount
func coinChange(coins []int, amount int) []int {
	bestValues := make([]int, amount+1)
	//sort coins here, if necessary

	//initialize starting coin values
	for _, value := range coins {
		if value < amount {
			bestValues[value] = value
		}
	}

	for index := 0; index <= amount; index++ {
		for descIndex := len(coins) - 1; descIndex >= 0; descIndex-- {
			if index-coins[descIndex] < 0 {
				continue
			} else {
				if bestValues[index-coins[descIndex]] > 0 {
					bestValues[index] = coins[descIndex]
					break
				}
			}
		}
	}

	result := []int{}

	for amount > 0 {
		result = append(result, bestValues[amount])
		amount -= bestValues[amount]
	}
	return result
}
