package main

import "fmt"

func main() {
	var choice, n1, n2, result int
	for {
		choice = getUserChoice()
		if choice == 5 {
			break
		}
		n1, n2 = getNumbers()
		result = processInputs(choice, n1, n2)
		fmt.Println("Result = ", result)
	}
}

func processInputs(choice, n1, n2 int) int {
	var result int
	switch choice {
	case 1:
		result = add(n1, n2)
	case 2:
		result = subtract(n1, n2)
	case 3:
		result = multiply(n1, n2)
	case 4:
		result = divide(n1, n2)
	}
	return result
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
