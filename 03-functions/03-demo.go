package main

import "fmt"

func main() {
	//function assigned to a variable
	var addFn func(int, int) int
	addFn = func(x, y int) int {
		return x + y
	}
	fmt.Println(addFn(100, 200))

	//anonymous functions
	func() {
		fmt.Println("anonymous fn invoked")
	}()

	func(x, y int) {
		fmt.Println("Subtract y from x (anonymous function) = ", x-y)
	}(100, 200)

	//functions as arguments
	compute(add, 100, 200)
	compute(subtract, 1000, 2000)

	adder := getAdder()
	fmt.Println(adder(50, 20))

	adderFor100 := getAdderFor(100)
	fmt.Println(adderFor100(200))
}

func getAdderFor(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func getAdder() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func compute(fn func(int, int) int, x, y int) {
	fmt.Printf("Computing %d and %d, result = %d\n", x, y, fn(x, y))
}
