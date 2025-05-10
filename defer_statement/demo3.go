package deferstatement

import "fmt"

func Log(functionName string, request product) {
	fmt.Println(functionName, request)
}

type product struct {
	name  string
	price float64
}

func (p product) save_product() {
	defer Log("save_product", p)

	if p.price <= 0 {
		panic("Something went wrong")
	}
	fmt.Println("Ürün kaydedildi:", p.name)
}

func Demo3() {
	p := product{name: "IPhone 11", price: 45.234}
	p.save_product()
	defer p.save_product()
	p = product{name: "Nokia", price: -324} // defer bunu değil kendinden öncekini alır
	defer p.save_product()                  // burada Nokia olacaktır
}
