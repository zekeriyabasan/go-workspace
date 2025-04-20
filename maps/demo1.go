package maps

import "fmt"

var numbers = map[int]string{1: "bir", 2: "iki", 3: "üç", 4: "dört"}

func FetchDictionaryValue(key string) (string, bool) {

	colors := make(map[string]string)
	colors["blue"] = "mavi"
	colors["yellow"] = "sarı"
	colors["black"] = "siyah"

	value, exist := colors[key]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Böyle bir kayıt yok !")
	}

	return value, exist

}
func FetchNumberText(number int) (string, bool) {

	fmt.Println(numbers)

	value, exist := numbers[number]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Böyle bir kayıt yok !")
	}
	return value, exist
}
func DeleteANumber(number int) bool {
	fmt.Println(numbers)

	_, exist := numbers[number]

	if exist {
		delete(numbers, number)
		fmt.Println("Silindikten sonra: ", numbers)
		return true
	} else {
		fmt.Println("Böyle bir kayıt yok !")
	}

	return false
}
func AddANumber(number int, text string) bool {
	_, exist := numbers[number]
	if exist {
		fmt.Println("Bu kayıt zaten var !")
		return false
	} else {
		numbers[number] = text
		fmt.Println(numbers)
		return true
	}

}
func UpdateANumber(number int, text string) bool {
	_, exist := numbers[number]
	if exist {
		numbers[number] = text
		fmt.Println(numbers)
		return true
	} else {
		fmt.Println("Böyle bir kayıt yok !")
		return false
	}
}
