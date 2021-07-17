package main

import "fmt"

func main() {
	// Array
	var ages [3]int = [3]int{20, 25, 30} // array with integer type and 3 length
	var agesTwo = [3]int{20, 25, 30}     // same with above
	fmt.Println(ages, agesTwo)
	fmt.Println("array length is: ", len(ages)) // two print array length

	// shorthand to create an array
	names := [4]string{"ashandi", "lutpi", "eman", "palah"}
	names[1] = "adam"
	fmt.Println(names, len(names))

	// Slices (use arrays under the hood)
	// slice can append
	var scores = []int{100, 50, 30}
	scores[2] = 25
	scores = append(scores, 85)
	fmt.Println("Slice: ", scores)

	// Slice ranges -> way to get a range of elements
	rangeOne := names[1:]
	fmt.Println(rangeOne)
}
