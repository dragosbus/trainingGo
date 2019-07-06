package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		y := 40
		fmt.Println(y)
	}
	//fmt.Println(y)--outside the scope of y
	fmt.Println(z)
	increment()
	fmt.Println(z)

	increment := wrapper()
	fmt.Println(increment())
}

var z int

func increment() int {
	z++
	return z
}

//Return function

func wrapper() func() int {
	y := 0
	return func() int {
		y++
		return y
	}
}
