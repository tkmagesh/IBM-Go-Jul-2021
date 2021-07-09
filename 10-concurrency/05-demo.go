package main

import (
	"fmt"
	"sync"
	"time"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {
	count := make(chan int)
	wg.Add(1)
	go fn(count)
	fmt.Println("Writing the count")
	time.Sleep(time.Second * 4)
	count <- 10
	wg.Wait()
	fmt.Println("Exiting from the main")
}

func fn(count chan int) {
	fmt.Println("Attempting to read the count")
	max := <-count
	fmt.Println("Just read the data from the count channel")
	for i := 0; i < max; i++ {
		fmt.Println(i)
	}
	wg.Done()
}
