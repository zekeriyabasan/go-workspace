package structs

import "fmt"

func Demo1() {
	fmt.Println(product{id: 1, name: "Computer ASUS K555Ls", unitPrice: 888.55, brand: "ASUS"})

}

type product struct {
	id        int
	name      string
	unitPrice float32
	brand     string
}
