package functions

func MultiReturnAddRemoveDivide(a int, b int) (int, int, float32) {
	added := a + b
	removed := a - b
	divided := float32(a / b)

	return added, removed, divided

}
