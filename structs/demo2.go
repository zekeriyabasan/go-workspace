package structs

import "fmt"

var customers = []Customer{}

func (c Customer) Add() {
	customers = append(customers, c)
	fmt.Println(201)
}
func (c Customer) Update(id int) {
	for i := range customers {
		if customers[i].Id == c.Id {
			customers[i] = c
			fmt.Println("Customer updated.")
			return
		}
	}
	fmt.Println("Customer not found.")
}

func GetCustomers() []Customer {
	return customers
}

func AddACustomer(c Customer) {
	c.Add()
}
func UpdateACustomer(c Customer) {
	c.Update(c.Id)
}
