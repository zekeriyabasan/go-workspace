package goroutines

import (
	"fmt"
	"time"
)

func WriteColors(colors []string) {
	for _, color := range colors {
		fmt.Println(color)
		time.Sleep(2 * time.Second)
	}

}
func WriteNumbers(numbers []int) {
	for _, num := range numbers {
		fmt.Println(num)
		time.Sleep(1 * time.Second)
	}

}
