package for_range

import (
	"fmt"
	"slices"
)

var oddNumbers = []int{23, 53, 1, 77, 87}

func AddOddNumbers(numbers ...int) int {
	var result = 0
	for i, number := range numbers {

		if number%2 != 0 {
			result = result + number
			AddOddNumberToOdd(number)
		}
		println("index: ", i)
	}
	return result
}

func AddOddNumberToOdd(number int) {
	if !slices.Contains(oddNumbers, number) {
		oddNumbers = append(oddNumbers, number)
		fmt.Println("ODDS: ", oddNumbers)
	}

}

func WriteOddNumbers() {
	for _, oddNumber := range oddNumbers {
		fmt.Println(oddNumber)
	}
}

func WriteToConsoleMap(colors map[string]string, seperator string) {
	for key, value := range colors {
		fmt.Println(key + seperator + value)
	}
}
