package main

import "fmt"

func main() {
	fmt.Println(divideClient(100, 0))
}

func divideClient(x, y int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Something went wrong!")
			fmt.Println(r)
		} else {
			fmt.Println("All is well!")
		}
	}()
	return div(x, y)
}

func div(x, y int) int {
	return divide(x, y)
}

func divide(x, y int) int {
	if y == 0 {
		panic("Invalid arguments!")
	}
	return x / y
}
