package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func main() {
	//var p Product
	//var p Product = Product{102, "Pen", 10, 20, "Stationary"}

	var p Product = Product{Id: 100, Name: "Pen", Cost: 10, Units: 20, Category: "Stationary"}
	fmt.Println(p)
	fmt.Println("Cost of Pen => ", p.Cost)

	/* productPtr := &Product{Id: 100, Name: "Pen", Cost: 10, Units: 20, Category: "Stationary"}
	//fmt.Println("Cost of pen => ", (*productPtr).Cost)
	fmt.Println("Cost of Pen => ", productPtr.Cost)
	fmt.Println(productPtr) */

	applyDiscount(&p, 10)
	fmt.Println("Product cost after discount => ", p.Cost)
}

func applyDiscount(product *Product, discount int) {
	product.Cost = product.Cost * ((100 - float32(discount)) / 100)
}
