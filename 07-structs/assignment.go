package main

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
		Product{104, "Scribble Pad", 20, 20, "Stationary"}
	}

}

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
