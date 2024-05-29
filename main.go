package main

import "fmt"

func main() {
	const age = 33
	fmt.Printf(" current age: %v %T", age, age)

	var str_lit = `this long string literal
	
	asdfasdfasdfasdf
	asdfasdfsadfsad
	`

	fmt.Printf("%v", str_lit)

	this_arr := [3]int{1, 2, 3}
	for _, val := range this_arr {
		fmt.Println(val)
	}

	events := map[string]string{"msuic": "heradsfadf"}
	fmt.Println(events["msuic"])

	if _, ok := events["msuic"]; ok {
		fmt.Println("event exits")
		fmt.Println(ok)
	} else {
		fmt.Println("event does not exist")
	}
}
