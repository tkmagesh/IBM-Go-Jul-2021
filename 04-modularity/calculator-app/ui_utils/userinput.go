package ui_utils

import (
	"fmt"

	"github.com/fatih/color"
)

func GetNumbers() (int, int) {
	var n1, n2 int
	color.Yellow("Enter number 1 :")
	fmt.Scanf("%d\n", &n1)
	color.Yellow("Enter number 2 :")
	fmt.Scanf("%d\n", &n2)
	return n1, n2
}

func GetUserChoice() int {
	var choice int
	color.Cyan("Enter the choice :")
	color.Cyan("1 : Add")
	color.Cyan("2 : Subtract")
	color.Cyan("3 : Multiply")
	color.Cyan("4 : Divide")
	color.Cyan("5 : Exit")
	fmt.Scanf("%d\n", &choice)
	return choice
}
