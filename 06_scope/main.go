package main

import "fmt"

//package level scope
var x = 42

func main() {
	fmt.Println(x)
	foo()
	//block level scope
	y := 55
	fmt.Println(y)
	callMax()
}

func foo() {
	fmt.Println(x)
	fmt.Println(z)
}

//for vars enclosed in package level scope, the order is dont matter

var z = 33

//dont ever do shadowing
func max(x int) int {
	return 40 + x
}

func callMax() {
	//DONT.If the max is called after the reasignment, the max is the value and not the function
	max := max(7)
	fmt.Println(max)
}
