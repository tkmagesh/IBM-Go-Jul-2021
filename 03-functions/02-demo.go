package main

import "fmt"

func main() {
	fmt.Println(add(10, 20))
	fmt.Println(add(10, 20, 30, 40))
	fmt.Println(add(10, 20, 30, 40, 50))
}

func add(nos ...int) int {
	//nos => slice (collection), nos[0], nos[1]..., len(nos)
	result := 0
	/*
		for index := 0; index < len(nos); index++ {
			result += nos[index]
		}
	*/
	for _, no := range nos {
		result += no
	}
	return result
}
