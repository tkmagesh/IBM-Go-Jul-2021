package main

import (
	"math"
	"fmt"
)

type Circle struct{
	Radius float32
}

func (c Circle) Area() float32{
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Height float32
	Width float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}


//utility function
func PrintArea( ){
}

func main(){
	c := Circle{Radius : 10}
	fmt.Println("Area = ", c.Area())

	r := Rectangle{ Height : 10, Width : 12 }
	fmt.Println("Area = ", r.Area())
}