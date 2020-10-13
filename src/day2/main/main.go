package main

import (
	"day2/domain"
	"day2/mapstore"
)

func main() {
	controller := CustomController{
		store: mapstore.NewMapStore(),
	}
	customer := domain.Customer{
		ID:   "01",
		Name: "JFM",
	}
	controller.Add(customer)
	controller.GetAll()
	controller.GetById("01")
	controller.GetById("02")
}
