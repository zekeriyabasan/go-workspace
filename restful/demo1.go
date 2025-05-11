package restful

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func Demo1() {
	jsonPlaceHolderBaseUrl := "https://jsonplaceholder.typicode.com/"
	response, err := http.Get(jsonPlaceHolderBaseUrl + "todos")

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	byteResult, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var todos []Todo
	err = json.Unmarshal(byteResult, &todos)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Todos: %+v\n", todos)
}
