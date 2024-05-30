package main

import "fmt"

func main() {
	fmt.Println("fundamentals")

	/*
		variables

		- Use variables if declared or else its considerred code pollution
		- use short form declaration unless there is a zero value declaratiob
	*/

	// Variables
	var a int = 22
	fmt.Printf("variable declaration: %v\n", a)

	b := 11
	fmt.Printf("Short declaration: %v\n", b)

	// inline declaration
	c, d, e := 5, 6, 7
	fmt.Printf("inline declaration: %v %v %v \n", c, d, e)

	// zero value
	var f int // declared but value not assigned
	fmt.Printf("zero value: %v\n", f)

	// printing values as binary & hexadecimal
	g := 8
	fmt.Printf("binary: %b  hex %#X\n", g, g)

	/*
		Types
		conversation
	*/

	// convert float 32 to float 64
	// considered float64
	z := 55.00
	fmt.Printf("z: %v type %T \n", z, z)

	var h float32 = 55.111

	z = float64(h)
	fmt.Printf("converting h float32 to 64 z: %v type %T \n", z, z)
	fmt.Printf("converting z float64 to int %v  Type: %T\n", int(z), int(z))

	/*
		Functions
	*/

	sayHello()
	result := add(2, 4)
	fmt.Printf("sum of add func %v \n", result)

}

func sayHello() {
	fmt.Println("sayHello func call \n ")
}

func add(x int, y int) int {
	/*
		erturn the sum of both ints passed as args

		return type: int
	*/
	return x + y
}
