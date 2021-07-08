package main

import (
	"fmt"
)

var operations map[int]func(int, int) int = map[int]func(int, int) int{
	1: add,
	2: subtract,
	3: multiply,
	4: divide,
}

func main() {
	var choice, n1, n2 int
	for {
		choice = getUserChoice()
		if _, exists := operations[choice]; !exists {
			break
		}
		n1, n2 = getNumbers()
		result := processInputs(choice, n1, n2)
		fmt.Println("Result = ", result)
	}
}

func processInputs(choice, n1, n2 int) int {
	oper := operations[choice]
	return oper(n1, n2)
}

func getNumbers() (int, int) {
	var n1, n2 int
	fmt.Println("Enter number 1 :")
	fmt.Scanf("%d\n", &n1)
	fmt.Println("Enter number 2 :")
	fmt.Scanf("%d\n", &n2)
	return n1, n2
}

func getUserChoice() int {
	var choice int
	fmt.Println("Enter the choice :")
	fmt.Println("1 : Add")
	fmt.Println("2 : Subtract")
	fmt.Println("3 : Multiply")
	fmt.Println("4 : Divide")
	fmt.Println("5 : Exit")
	fmt.Scanf("%d\n", &choice)
	return choice
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
