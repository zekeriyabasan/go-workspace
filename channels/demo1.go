package channels

import (
	"fmt"
)

func SumOddNumber(sumOddCn chan int, numbers []int) {
	var total = 0
	for _, num := range numbers {
		fmt.Println(num)
		if num%2 != 0 {
			total += num
		}
	}

	sumOddCn <- total // bu async işlemin bittiğini garanti ediyor channel a int dönüşü

}

func WriteColors(writedCn chan bool, colors []string) {
	for _, color := range colors {
		fmt.Println(color)
	}

	writedCn <- true

}
