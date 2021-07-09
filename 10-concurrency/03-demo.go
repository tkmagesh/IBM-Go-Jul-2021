package main

import (
	"fmt"
	"sync"
)

//Communicate by sharing memory
var result int

var wg *sync.WaitGroup = &sync.WaitGroup{}
var mutex *sync.Mutex = &sync.Mutex{}

func main() {
	wg.Add(1)
	go add(100, 200)
	wg.Wait()
	fmt.Println(result)
}

func add(x, y int) {
	mutex.Lock()
	result = x + y
	mutex.Unlock()
	wg.Done()
}
