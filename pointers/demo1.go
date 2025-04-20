package pointers

import (
	"fmt"
)

func Demo1(number *int) {
	*number = *number + 1
	fmt.Println("DEMO : ", *number)

}
func Demo2(numbers []int) {
	for _, num := range numbers {
		num = num + 1
	}
	fmt.Println("DEMO : ", numbers)
}
