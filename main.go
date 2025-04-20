package main

import (
	"everymore-go/functions"
	"fmt"
)

func main() {
	//variables.FirstDemo()
	//conditionals.CreateAccounting(12314.33, 41414)
	//conditionals.Demo1()
	//loops.PredicateGame()
	//loops.WriteFriendshipNumber(17296, 1846)
	//slices.Demo1()
	//functions.HiFunction("Zekeriya BAŞAN")
	//functions.Add(134, 400)
	added, removed, divided := functions.MultiReturnAddRemoveDivide(56, 7)

	fmt.Println("toplam: ", added)
	fmt.Println("fark: ", removed)
	fmt.Println("bölüm: ", divided)

	_, difference, _ := functions.MultiReturnAddRemoveDivide(56, 7)
	fmt.Println("Sadece farkı aldık :", difference)

}
