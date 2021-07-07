package main

import "fmt"

func main() {
	/* defer func() {
		fmt.Println("deferred from main")
	}() */
	defer fmt.Println("deferred from main")
	fmt.Println("main started")
	f1()
	fmt.Println("main completed")
}

func f1() {
	defer fmt.Println("[defer - f11] deferred from f1")
	defer fmt.Println("[defer - f12] deferred from f1")
	defer fmt.Println("[defer - f13] deferred from f1")
	fmt.Println("f1 invoked")
	msg := f2()
	fmt.Println(msg)
}

func f2() (message string) {
	defer func() {
		fmt.Println("[defer - f21] deferred from f2")
		message = "Message from the deferred f21 fn"
	}()
	defer fmt.Println("[defer - f22] deferred from f2")
	defer fmt.Println("[defer - f23] deferred from f2")
	fmt.Println("f2 invoked")
	message = "message from the f2 func"
	return
}
