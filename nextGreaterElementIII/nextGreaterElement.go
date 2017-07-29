package main

import (
	"fmt"
	"strconv"
)

func main() {
	testValue := 130283122
	fmt.Println(testValue)
	result := nextGreaterElement(testValue)
	fmt.Println(result)
}

func nextGreaterElement(startValue int) int64 {
	stringArray := []rune(getSingleDigitArray(startValue))
	var bucket [10]int
	result := int64(0)
	for reverseIndex := len(stringArray) - 1; reverseIndex >= 0; reverseIndex-- {

		value := stringArray[reverseIndex]
		higherElementIndex := bucketContainsHigherElement(bucket, int(value))
		if higherElementIndex > 0 {
			stringArray[higherElementIndex], stringArray[reverseIndex] = stringArray[reverseIndex], stringArray[higherElementIndex]
			currentIndex := reverseIndex

			temp := mySort(stringArray[currentIndex+1:])
			stringArray = stringArray[:currentIndex+1]
			stringArray = append(stringArray, temp...)
			temp2, _ := strconv.ParseInt(string(stringArray), 10, 32)
			return temp2
		} else {
			//bucket value at that integer is the correct index for it
			if bucket[value-48] == 0 {
				bucket[value-48] = reverseIndex
			}
		}

	}

	return result
}

func mySort(array []rune) []rune {
	if len(array) < 2 {
		return array
	}

	mid := len(array) / 2
	return merge(mySort(array[:mid]), mySort(array[mid:]))
}

func merge(left, right []rune) []rune {
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]rune, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

func bucketContainsHigherElement(bucket [10]int, value int) int {
	for index := value - 48 + 1; index < 10; index++ {
		if bucket[index] > 0 {
			return bucket[index]
		}
	}
	return 0
}

//returns an integer array where position is the decimal digit of original, and value is that digit
func getSingleDigitArray(value int) string {
	stringValue := strconv.Itoa(value)
	return stringValue
}
