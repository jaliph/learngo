package main

import (
	"day2/domain"
	"day2/mapstore"
	"fmt"
)

type CustomController struct {
	store *mapstore.MapStore
}

func (c CustomController) Add(cus domain.Customer) {
	err := c.store.Create(cus)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("New Customer has been created")
}
func (c CustomController) Delete(s string) {
	err := c.store.Delete(s)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Customer Deleted!")
}
func (c CustomController) Update(s string, cus domain.Customer) {
	err := c.store.Update(s, cus)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Customer has been updated")
}
func (c CustomController) GetAll() {
	val, err := c.store.GetAll()
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("All Customers are")
	fmt.Println(val)
}
func (c CustomController) GetById(s string) {
	val, err := c.store.GetByID(s)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Customers found by Id")
	fmt.Println(val)
}

func main() {
	controller := CustomController{
		store: mapstore.NewMapStore(),
	}
	customer := domain.Customer{
		ID:   "01",
		Name: "JFM",
	}
	controller.Add(customer)
}
