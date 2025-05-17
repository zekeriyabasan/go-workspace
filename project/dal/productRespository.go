package dal

import (
	"everymore-go/project/dal/base/concrete"
	entities "everymore-go/project/models"
	"fmt"
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
