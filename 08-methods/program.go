package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (product *Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s\n", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

//implement the applyDiscount "method" for the product

func main() {
	pencil := Product{107, "Pencil", 2, 100, "Stationary"}
	//fmt.Print(Format(pencil))
	fmt.Print(pencil.Format())

}
