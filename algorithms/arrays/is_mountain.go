/*
Given an array determine if its a valid mountain.

It is a mountain if the array contains 3 elements or more
it is a mountain if the array starts increasing then has a subsequence that decreases

valid = [1,2,3,2]
invalid = [1]
invalid = [1,2,3,3]

first check if there are 3 or more elemenets

- iterate through the array to go through increasing sequenec and stop once you get to decreasing
- check that you havent went through the entire arr or that you have even moved through the array
- iterate through the decreasing sequence

*/

package main

import "fmt"

func is_mountain(arr [4]int) bool {
	fmt.Println(arr)

	// base case
	if len(arr) < 3 {
		return false
	}
	// start at second element since we are comparing previous and current
	L := 1

	// check for increasing order
	for L < len(arr) && arr[L-1] < arr[L] {
		L++
	}

	// exit if didnt move or moved through entire array
	if L == 1 || L == len(arr) {
		return false
	}

	// check ecreasing order
	for L < len(arr) && arr[L-1] > arr[L] {
		L++
	}

	// return as true if L gets to the end of the array
	return L == len(arr)

}

func main() {
	fmt.Println("is valid mountain array")

	mountain := [4]int{1, 2, 3, 2}
	fmt.Println(is_mountain(mountain))

}
