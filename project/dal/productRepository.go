package productDdal

import (
	"encoding/json"
	config "everymore-go/project"
	"everymore-go/project/entities"
	"fmt"
	"io"
	"net/http"
)

func GetAll() {
	response, err := http.Get(config.JsonServerBaseAdress + "products")

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	byteResult, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var products []entities.Product
	err = json.Unmarshal(byteResult, &products)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Products: %+v\n", products)
}
