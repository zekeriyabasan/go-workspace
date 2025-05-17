package dal

import (
	"bytes"
	"encoding/json"
	config "everymore-go/project"
	"everymore-go/project/dal/base/concrete"
	entities "everymore-go/project/models"
	"fmt"
	"net/http"
)

func AddProduct(repo *concrete.JsonRepository[*entities.Product], product *entities.Product) error {
	err := repo.Add(product)
	if err != nil {
		fmt.Println("Ürün eklenemedi:", err)
		return err
	}
	fmt.Println("Ürün başarıyla eklendi:", product.Name)
	return nil
}
func GetProducts(repo *concrete.JsonRepository[*entities.Product]) ([]*entities.Product, error) {
	products, err := repo.GetAll()
	if err != nil {
		fmt.Println("Ürünler listelenemedi:", err)
		return nil, err
	}

	fmt.Println("Ürünler listesi:")
	for _, p := range products {
		fmt.Printf("- ID: %d | Name: %s\n", p.GetID(), p.Name)
	}

	return products, nil
}

func AddProduct2() {
	product := entities.Product{Name: "Subrosa", CategoryID: 2, UnitPrice: 44323.234, Views: 123}
	data, err := json.Marshal(product)

	if err != nil {
		fmt.Println(err)
	}
	response, err := http.Post(config.JsonServerBaseAdress+"products", "application/json;charset:utf-8", bytes.NewBuffer(data))

	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	fmt.Println("Kaydedildi")
}
