package main

import "fmt"

func main() {
	age := 25

	if age < 30 {
		fmt.Println("age is later than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 45")
	}

	names := []string{"ashandi", "adam", "dani", "afalah", "rochim", "alip", "iqbale"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue // continue to the loop
		}

		if index > 2 {
			fmt.Println("breaking at post", index)
			break // break the loop and dont continue the loop
		}

		fmt.Printf("the value at pos %v is %v \n", index, value)
	}
}
