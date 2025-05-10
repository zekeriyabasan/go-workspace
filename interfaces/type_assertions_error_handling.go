package interfaces

import "fmt"

func validateInt(value interface{}) {
	num, ok := value.(int)
	fmt.Println(num, ok)
}

func Demo3() {
	validateInt(5)
	validateInt("zek")

	var num interface{} = 10
	validateInt(num)

	var name interface{} = "Zekeriya"
	validateInt(name)

}
