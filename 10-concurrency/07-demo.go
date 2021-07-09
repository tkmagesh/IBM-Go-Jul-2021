/*
modify the program to print the following
Hello
World
Hello
World
Hello
World
Hello
World
Hello
World
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	go print("Hello", 1*time.Second)
	go print("World", 3*time.Second)
	var input string
	fmt.Scanln(&input)
}

func print(s string, delay time.Duration) {
	for i := 0; i < 5; i++ {
		println(s)
		time.Sleep(delay)
	}
}
