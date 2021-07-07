package main

import "fmt"

func main() {
	loggerAdd := getLoggerFn(add, "add")
	fmt.Println(loggerAdd(100, 200))
	loggerSubtract := getLoggerFn(subtract, "subtract")
	fmt.Println(loggerSubtract(100, 200))
}

func getLoggerFn(fn func(int, int) int, operName string) func(int, int) int {
	return func(x, y int) int {
		fmt.Printf("[%s] Beginning invocation\n", operName)
		result := fn(x, y)
		fmt.Println("Result = ", result)
		fmt.Printf("[%s] After invocation\n", operName)
		return result
	}
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

/*
	create a logger version of these functions which will print the following
	for ex, invoking the add/subtract function should display as below:

		Begin invocation
		300
		After invocation

	NOTE : treat the add & subtract functions as NOT CHANGABLE
*/
