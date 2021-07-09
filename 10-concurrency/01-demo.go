package main

import (
	"fmt"
	"time"
)

func main() {
	go f1()
	go f2()
	/*
		var input string
		fmt.Scanln(&input)
	*/
	time.Sleep(time.Second * 1)
}

func f1() {
	fmt.Println("f1 is invoked")
}

func f2() {
	fmt.Println("f2 is invoked")
}
