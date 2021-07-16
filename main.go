package main

import "fmt"

func main() {
	// Numbers

	// use int to declared integer
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40 // you can also use this shorthand to create int variable

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory -> full documentation https://pkg.go.dev/builtin#int
	var numOne int8 = 25 // means 8 bits with Range: -128 through 127.
	var numTwo int8 = -128
	var numThree uint8 = 23 // we can't have a negative number

	fmt.Println(numOne, numTwo, numThree)

	// float -> decimal number
	var scoreOne float32 = 25.98
	scoreTwo := 291292.21 // if we dong specify the type of float number, the type will be float64 as default

	fmt.Println(scoreOne, scoreTwo)
}
