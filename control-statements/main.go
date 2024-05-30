package main

import "fmt"

func main() {

	/* If else */
	decisionMaking()
	// use as while loop
	whileLoop()

	this_arr := [5]int{5, 4, 7, 88, 2}
	// for loops
	traditionalLoop(this_arr)
	rangeLoop(this_arr)

	// switch
	switchStatement()

}

func switchStatement() {
	fmt.Print("switchStatement\n")
	this_string := "five"

	switch this_string {
	case "one":
		fmt.Println("Monday")
	case "two":
		fmt.Println("Tuesday")
	case "three":
		fmt.Println("Wednesday")
	case "four":
		fmt.Println("Thursday")
	case "five":
		fmt.Println("Friday")
	}
}

func traditionalLoop(arr [5]int) {

	fmt.Println("traditionalLoop")

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%v \n", arr[i])
	}

}

func rangeLoop(arr [5]int) {

	fmt.Println("rangeLoop")

	fmt.Println("print index and value")

	for x, y := range arr {
		fmt.Printf("index: %v val: %v \n", x, y)
	}

	fmt.Println("Print arr using the index")
	for i := range arr {
		fmt.Printf("%v \n", arr[i])
	}
}

// for loop as a while loop
func whileLoop() {
	x := 0

	for x < 5 {
		fmt.Printf("printing until x reaches case x: %v \n", x)
		x += 1
	}
}

/* If else */
func decisionMaking() {
	x := 100
	if x < 100 {
		fmt.Println("Value is less than 100")
	} else if x > 100 {
		fmt.Println("Value is greater")
	} else {
		fmt.Println("Value doesnt match any case")
	}

}
