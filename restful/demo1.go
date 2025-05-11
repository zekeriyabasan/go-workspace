package restful

import (
	"bytes"
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

const jsonPlaceHolderBaseUrl = "https://jsonplaceholder.typicode.com/"

func Demo1() {
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

func AddTodo() {

	todo := Todo{1, 2, "Tereyağı alınacak!", false}

	byteTodo, err := json.Marshal(todo)
	if err != nil {
		panic(err)
	}
	response, err := http.Post(jsonPlaceHolderBaseUrl+"todos", "application/json;charset:utf-8", bytes.NewBuffer(byteTodo))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	// Deserialize gibi
	byteResult, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var resultTodo Todo
	err = json.Unmarshal(byteResult, &resultTodo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Todo: %+v\n", resultTodo)

}
