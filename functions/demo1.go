package functions

import "fmt"

func Add(claim int, balance int) int {

	var newBalance = claim + balance
	fmt.Println("BALANCE: ", newBalance)

	return newBalance

}

func HiFunction(name string) string {

	fmt.Println("Hi : ", name)

	return "Hi," + name
}
