package main

import "fmt"

func main() {
	testCoins := []int{1, 4, 5}
	testAmount := 12

	result := coinChange(testCoins, testAmount)
	fmt.Println("testCoins: ", testCoins, "\ntestAmount: ", testAmount, "\nresult", result)
}

type Node struct {
	value int
	count int
}

//result is an array of coin values that most efficiently add up to  amount
func coinChange(coins []int, amount int) []int {
	bestValues := make([]Node, amount+1)
	//sort coins here, if necessary

	//initialize starting coin values
	for _, value := range coins {
		if value < amount {
			bestValues[value].value = value
			bestValues[value].count = 1
		}
	}
	fmt.Println("pre best values: ", bestValues, "\nend")
	for index := 0; index <= amount; index++ {
		bestCount := amount
		bestValue := 0
		for _, coinValue := range coins {
			if coinValue == index {
				bestCount = 0
				bestValue = coinValue
				break
			}
			if index-coinValue > 0 {
				temp := bestValues[index-coinValue].count
				if temp < bestCount {
					bestValue = coinValue
					bestCount = temp
				}
			}

		}

		bestValues[index].value = bestValue
		bestValues[index].count = bestCount + 1
	}
	for index, value := range bestValues {
		fmt.Println("index: ", index, "	value: ", value.value, "	count: ", value.count)
	}
	result := []int{}

	for amount > 0 {
		result = append(result, bestValues[amount].value)
		amount -= bestValues[amount].value
	}
	return result
}
