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
	ch1 := make(chan string)
	ch2 := make(chan string)
	go print("Hello", 1*time.Second, ch1, ch2)
	go print("World", 3*time.Second, ch2, ch1)
	ch2 <- "start"
	var input string
	fmt.Scanln(&input)
}

func print(s string, delay time.Duration, x chan string, y chan string) {
	for i := 0; i < 5; i++ {
		<-x
		println(s)
		time.Sleep(delay)
		y <- "done"
	}
}
