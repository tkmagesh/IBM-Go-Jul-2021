/*
Assignment
Write an add function to add 2 numbers

concurrently execute the add function to add 2 set of values (10,20) and (100,200)
aggregate them in the main function and print the final result (330)

*/
package main

import "fmt"

func main() {
	ch := make(chan int)
	go add(10, 20, ch)
	go add(100, 200, ch)
	result1 := <-ch
	result2 := <-ch
	go add(result1, result2, ch)
	finalResult := <-ch
	fmt.Println(finalResult)
}

func add(x, y int, ch chan int) {
	ch <- x + y
}
