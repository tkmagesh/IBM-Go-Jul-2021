package main

import (
	"calculator"
	uiutils "calculator-app/ui_utils"
	"calculator/utils"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println(utils.IsEven(100))
	var choice, n1, n2, result int
	for {
		choice = uiutils.GetUserChoice()
		if choice == 5 {
			break
		}
		n1, n2 = uiutils.GetNumbers()
		result = processInputs(choice, n1, n2)
		color.Green(fmt.Sprintf("Result = %d\n", result))
	}
}

func processInputs(choice, n1, n2 int) int {
	var result int
	switch choice {
	case 1:
		result = calculator.Add(n1, n2)
	case 2:
		result = calculator.Subtract(n1, n2)
	case 3:
		result = calculator.Multiply(n1, n2)
	case 4:
		result = calculator.Divide(n1, n2)
	}
	return result
}
