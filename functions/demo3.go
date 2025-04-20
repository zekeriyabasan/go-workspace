package functions

func AddVariadic(numbers ...int) int {
	var total = 0

	for i := 0; i < len(numbers); i++ {
		total = total + numbers[i]
	}

	return total

}
