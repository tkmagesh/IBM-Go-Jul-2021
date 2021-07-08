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
	products := []Product{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Stationary"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}

	dummyProduct := Product{Id: -100}
	pencil := Product{107, "Pencil", 2, 100, "Stationary"}
	fmt.Println("Index of Pencil => ", IndexOf(&products, pencil))
	fmt.Println("Includes Pencil ? => ", Includes(&products, pencil))
	fmt.Println("Includes dummyProduct ? => ", Includes(&products, dummyProduct))

	/*
		fmt.Println("Any costly products ? => ", AnyCostlyProduct(&products))
		fmt.Println("Any stationary products ? => ", AnyStationaryProduct(&products))
	*/
	areThereCostlyProducts := Any(&products, func(p Product) bool {
		return p.Cost > 50
	})
	fmt.Println("Are there any costly products ? => ", areThereCostlyProducts)

	stationaryProductCriteria := func(p Product) bool {
		return p.Category == "Stationary"
	}

	areThereStationaryProducts := Any(&products, stationaryProductCriteria)
	fmt.Println("Are there any stationary products ? => ", areThereStationaryProducts)

	areThereElectronicsProducts := Any(&products, func(p Product) bool {
		return p.Category == "Electronics"
	})
	fmt.Println("Are there any electronic products ? => ", areThereElectronicsProducts)

	areAllStationaryProducts := All(&products, stationaryProductCriteria)

	fmt.Println("Are all the products stationary products => ", areAllStationaryProducts)

	stationaryProducts := Filter(&products, stationaryProductCriteria)
	fmt.Println("Stationary Products")
	PrintList(stationaryProducts)
}

func Format(product *Product) string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s\n", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

func PrintList(products *[]Product) {
	for _, p := range *products {
		fmt.Print(Format(&p))
	}
}

func IndexOf(products *[]Product, product Product) int {
	for idx, p := range *products {
		if p == product {
			return idx
		}
	}
	return -1
}

func Includes(products *[]Product, product Product) bool {
	return IndexOf(products, product) != -1
}

func Any(products *[]Product, criteria func(Product) bool) bool {
	for _, product := range *products {
		if criteria(product) {
			return true
		}
	}
	return false
}

func All(products *[]Product, criteria func(Product) bool) bool {
	for _, product := range *products {
		if !criteria(product) {
			return false
		}
	}
	return true
}

func Filter(products *[]Product, criteria func(Product) bool) *[]Product {
	result := &[]Product{}
	for _, product := range *products {
		if criteria(product) {
			*result = append(*result, product)
		}
	}
	return result
}

/* func AnyCostlyProduct(products *[]Product) bool {
	for _, product := range *products {
		if product.Cost > 50 {
			return true
		}
	}
	return false
}

func AnyStationaryProduct(products *[]Product) bool {
	for _, product := range *products {
		if product.Category == "Stationary" {
			return true
		}
	}
	return false
} */

/*
write the following utility functions to manipulate the products
	IndexOf => returns the index of the given product
	Includes => return true/false based on the presence of the given product in the products list
	Any => return true /false based on the presence of atleast one product in the list based on the given criteria
		use case:
			Is there any cost product? (cost > 100)
			Is there any stationary product? (category == "stationary")
	All => return true /false based on the condition that all the products in the list satisfy the given criteria
		use case:
			Are all products cost products? (cost > 100)
			Are all products stationary products? (category == "stationary")
	Filter => filter the products based on the given criteria
		use case:
			filter utencils from the products list
			filter under stocked products from the products list (units < 10)
*/
