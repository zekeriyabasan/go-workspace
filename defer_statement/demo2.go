package deferstatement

import "fmt"

func Add(a float64, b float64) float64 {
	return (a + b)
}
func Remove(a float64, b float64) float64 {
	return (a - b)
}
func Divide(a float64, b float64) float64 {
	if b == 0.0 {
		panic("b cannot be zero")
	}
	return a / b
}

func Write_Console(a string) {
	fmt.Println(a)
}

func Calculate() {
	a := 1231.43
	// b := 4.0
	b := 0.0
	c := 123.223
	defer Write_Console(fmt.Sprintf("a: %.2f", a))
	defer Write_Console(fmt.Sprintf("b: %.2f", b))
	defer Write_Console(fmt.Sprintf("c: %.2f", c))

	bc := fmt.Sprintf("a: %.2f, b: %.2f, c: %.3f", a, b, c)
	defer Write_Console(bc)

	z := Divide(c, b)
	x := Add(a, c)
	y := Remove(b, a)

	fmt.Printf("Toplam: %.2f, Fark: %.2f, Bölüm: %.3f\n", x, y, z)

}
