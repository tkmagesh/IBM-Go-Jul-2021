package main

import "fmt"

//share memory by communicating

/*
//the below code leads to deadlock
var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(1)
	go add(100, 200, ch)
	wg.Wait()
	result := <-ch
	fmt.Println(result)
}

func add(x, y int, ch chan int) {
	result := x + y
	ch <- result
	wg.Done()
}
*/

func main() {
	ch := make(chan int)
	go add(100, 200, ch)
	result := <-ch
	fmt.Println(result)
}

func add(x, y int, ch chan int) {
	result := x + y
	ch <- result
}

/*
Writing data into the channel
ch <- data

Reading data from the channle
variable := <- ch
*/
