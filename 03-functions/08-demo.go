package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := divide(100, 7)
	if err == nil {
		fmt.Println("Result = ", result)
		return
	}
	fmt.Println(err)
}

func divide(x, y int) (result int, err error) {
	if y == 0 {
		err = errors.New("Invalid arguments")
		return
	}
	result = x / y
	return
}
