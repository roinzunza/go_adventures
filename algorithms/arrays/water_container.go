/*
Given n non-negative integers where each integer represetns the hiegh of a vertical line on a chat
- find two lines which together with x - axis forms a container, that holds the most amount of water
- return the area of that water

eg.
        0  1  2  3  4  5  6  7  8
arr = [ 1, 8, 6, 2, 5, 4, 8, 3, 7]

area = width * Length

length = difference of two indexes that will be the miiddle part of the container
width = the max height without overflowing the container

Left index = 1 value: 8
Right index = 8 value = 7

length = 8 - 1 = 7
width = 7 since left index is 8 and right is 7 we cannot exceed 7 or else it will overflow

*/

package main

import "fmt"

func minInt(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max_area(arr []int) int {

	L := 0
	R := len(arr) - 1
	max_area := 0

	// while loop
	for L < R {

		// AREA = LENGTH * WITDH
		length := R - L
		// smallest  arr[R]  between arr[L]
		width := minInt(arr[R], arr[L])
		this_area := length * width

		if this_area > max_area {
			max_area = this_area
		}

		// move one of the two pointers depending on which is currenty largest value

		if arr[R] > arr[L] {
			L++ // move Left pointer to the right
		} else {
			R-- //move the right pointer to the left
		}

	}

	return max_area

}

func main() {

	// two pointers L and R
	// L will start at the beginning & R at the end
	// get the AREA of the distance between L & R
	// Length is the distance between indexes or R - L
	// width is the minimum value of the value of R and L
	// to move through the array efficiently we want to stop where the largest element is between L and R
	// if arr[L] < arr[R] L ++ else R +

	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	max_area := max_area(arr)
	fmt.Printf("water container max area: %v", max_area)

}
