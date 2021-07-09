package main

import (
	"fmt"
	"sync"
	"time"
)

var wg *sync.WaitGroup = &sync.WaitGroup{}

func main() {

	wg.Add(1)
	go f1()

	wg.Add(1)
	go f2()

	wg.Wait()
	fmt.Println("Exiting from main")
}

func f1() {
	time.Sleep(5 * time.Second)
	fmt.Println("f1 is invoked")
	wg.Done()
}

func f2() {
	time.Sleep(3 * time.Second)
	fmt.Println("f2 is invoked")
	wg.Done()
}
