package conditionals

import "fmt"

func Demo1() {
	// n adet int değişkenin en büyük olanı yazdır
	numbers := []int{67, 44, 154, 10, 20, 195, 30, 40, 300}
	greateNum := 0

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > greateNum {
			greateNum = numbers[i]
			println(greateNum)
		}
	}
	fmt.Printf("En büyük sayi : %v", greateNum)

}
