package main

import "fmt"

func main() {
	/*
		counter := getCounter()
		fmt.Println(counter()) //=> 1
		fmt.Println(counter()) //=> 2
		fmt.Println(counter()) //=> 3
		fmt.Println(counter()) //=> 4
		fmt.Println(counter()) //=> 5
	*/

	increment, decrement := getCounters()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())

}

func getCounter() func() int { //=> step : 1
	var count int           //=> step : 2
	counter := func() int { //=> step : 3
		count++ //=> step : 4
		return count
	}
	return counter //=> step : 5
}

func getCounters() (func() int, func() int) {
	count := 0
	increment := func() int {
		count++
		return count
	}
	decrement := func() int {
		count--
		return count
	}
	return increment, decrement
}
