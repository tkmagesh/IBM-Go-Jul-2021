package main

import (
	"math"
	"fmt"
)

type Circle struct{
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}



type Rectangle struct {
	Height float32
	Width float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32{
	return (2 * r.Height) + (2 * r.Width)
}


//utility function

type ShapeWithArea interface {
	Area() float32
}

func PrintArea(sa ShapeWithArea){
	fmt.Println("Area = ", sa.Area())
}

type ShapeWithPerimeter interface {
	Perimeter() float32
}

func PrintPerimeter(sp ShapeWithPerimeter){
	fmt.Println("Perimeter = ", sp.Perimeter())
}


/* interface composition */

/* 
type Dimension interface{
	Area() float32
	Perimeter() float32
}

func printDimensions(d Dimension){
	fmt.Println("Area = ", d.Area())
	fmt.Println("Perimeter = ", d.Perimeter())
} 
*/

type Dimension interface{
	ShapeWithArea
	ShapeWithPerimeter
}

func PrintDimensions(d Dimension){
	PrintArea(d)
	PrintPerimeter(d)
} 

func main(){
	c := Circle{Radius : 10}
	/* 
	PrintArea(c)
	PrintPerimeter(c) 
	*/
	PrintDimensions(c)

	r := Rectangle{ Height : 10, Width : 12 }
	/* PrintArea(r)
	PrintPerimeter(r) */
	PrintDimensions(r)
}