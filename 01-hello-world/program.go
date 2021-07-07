//package declaration
package main

//import other packages
import "fmt"

//package level variables / types

//package init function

func main() {

	/*
		var message string
		message = "Hello World"
	*/

	/*
		var message string = "Hello World"
	*/

	/*
		var message = "Hello World"
	*/

	//the following syntax works only in a function
	message := "Hello World"
	fmt.Println(message)

	//using multiple variables
	/*
		var name string
		var msg string
		name = "Magesh"
		msg = "Welcome to Golang"
	*/
	/*
		var (
			name string
			msg  string
		)
	*/
	/*
		var (
			name, msg string
		)
	*/
	/*
		var name, msg string
		name = "Magesh"
		msg = "Welcome to Golang"
	*/

	/*
		var name, msg string = "Magesh", "Welcome to Golang"
	*/

	/*
		var name, msg = "Magesh", "Welcome to Golang"
	*/
	name, msg := "Magesh", "Welcome to Golang"
	fmt.Println(name, msg)
}

//other functions
