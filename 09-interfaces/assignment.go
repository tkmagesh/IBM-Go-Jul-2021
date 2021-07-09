package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

/* implementing the Stringer interface */
func (product Product) String() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %s\n", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

//products

type Products []Product

/* implementing the Stringer interface */
func (products Products) String() string {
	result := ""
	for _, p := range products {
		result += fmt.Sprintf("%s", p)
	}
	return result
}

/* implementing the 'interface' interface */
func (products Products) Len() int { 
	return len(products) 
}

func (products Products) Swap(i, j int) { 
	products[i], products[j] = products[j], products[i] 
}

func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

/* To sort by name */
type ByName struct {
	Products
}

func (products ByName) Less(i, j int) bool {
	return products.Products[i].Name < products.Products[j].Name
}

/* Products Methods */

func (products Products) Sort(attrName string){
	switch attrName {
	case "Id":
		sort.Sort(products)	
	case "Name":
		sort.Sort(ByName{products})
	default:
		sort.Sort(products)	
	}
}


func main() {
	
	//manipulating products

	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Stationary"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}

	fmt.Println("Initial List")
	fmt.Println(products)

	fmt.Println("Default Sort")
	products.Sort("Id")
	fmt.Println(products)

	fmt.Println("Sort by name")
	products.Sort("Name")
	fmt.Println(products)
	
}
