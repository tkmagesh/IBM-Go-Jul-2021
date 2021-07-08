package main

import "fmt"

func main() {
	var i int = 100
	var iPtr *int = &i
	fmt.Println("Address of i = ", iPtr)

	//deferencing
	fmt.Println("Value at address is ", *iPtr)

	fmt.Println("[By value] Before incrementing, i = ", i)
	i = increment(i)
	fmt.Println("[By value] After incrementing, i = ", i)

	fmt.Println("[By address] Before incrementing, i = ", i)
	incrementByPtr(&i)
	fmt.Println("[By address] After incrementing, i = ", i)

	x, y := 10, 20
	swap(&x, &y)
	fmt.Println("After swapping ", x, y)

	nos := []int{10, 20, 30}
	addValue(&nos, 40)
	fmt.Println(nos) // => should print 10,20,30,40
}

func addValue(nos *[]int, no int) {
	*nos = append(*nos, no)
}

func increment(no int) int {
	fmt.Println("In the increment func, address = ", &no)
	no = no + 1
	return no
}

func incrementByPtr(noPtr *int) {
	fmt.Println("In the incrementByPtr func, address = ", noPtr)
	*noPtr = *noPtr + 1
}

func swap(x, y *int) {
	*x, *y = *y, *x
}
