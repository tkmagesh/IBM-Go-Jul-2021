package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

/*
type PerishableProduct struct {
	Prod   Product
	Expiry string
}
*/

type PerishableProduct struct {
	Product
	Expiry string
}

func main() {
	//grapes := PerishableProduct{Prod: Product{Id: 200, Name: "Grapes", Cost: 70, Units: 10, Category: "Food"}, Expiry: "2 Days"}
	//grapes := PerishableProduct{Product{200, "Grapes", 70, 10, "Food"}, "2 Days"}
	//grapes := PerishableProduct{Product: Product{200, "Grapes", 70, 10, "Food"}, Expiry: "2 Days"}

	grapes := NewPerishableProduct(200, "Grapes", 70, 10, "Food", "2 Days")
	fmt.Println(grapes)
	fmt.Println("Cost of grapes => ", grapes.Cost)
}

func NewPerishableProduct(id int, name string, cost float32, units int, category string, expiry string) *PerishableProduct {
	return &PerishableProduct{Product{id, name, cost, units, category}, expiry}
}
