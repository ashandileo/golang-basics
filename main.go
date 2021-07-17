package main

import "fmt"

func main() {

	age := 19
	name := "Ashandi"

	// Print
	// doesn't add a new line, if you want to add a new line use the (\n)
	fmt.Print("Hello \n")
	fmt.Print("World \n")

	// Println
	// will add a new line automatically at the end
	fmt.Println("Hello")
	fmt.Println("new line")

	fmt.Println("my age is", age, "and my name is", name)

	// Printf (Formatted string) -> way to create a string with variables embedded inside it
	// %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name) // %v default
	fmt.Printf("my age is %q and my name is %q \n", age, name) // %q will add double qoute for string var
	fmt.Printf("age is of type %T \n", age)                    // get the types of variable
	fmt.Printf("your scored %f points! \n", 225.25)            // float
	fmt.Printf("your scored %0.2f points! \n", 225.25555)      // float with 2 decimal

	// Sprintf (save formated string)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the save string is: ", str)
}
