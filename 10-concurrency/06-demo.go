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
