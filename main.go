package main

// the package allows you to format basic strings, values, or anything and print them
import "fmt"

// you can initialize variable outside main func
var someName = "Ashandy"

// but you can't create variable outside main func using this short hand, this will get an error
// anotherSomeName := "Leonadi"

func main() {

	// strings
	// use double quote when store string variable in go
	var nameOne string = "ashanday"

	// golang will automatically set the type of variable to string, if we dont declared the type
	var nameTwo = "lutpi"

	// initial variable with no value
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	// you can replace the variable value using this way
	nameOne = "anjay"
	nameThree = "coba ah ganti nama bisa gak"

	fmt.Println(nameOne, nameTwo, nameThree)

	// another way to initialize variable without var keyword by using colon (:)
	nameFour := "mantap"

	fmt.Println(nameFour)

	fmt.Println("variable dari luar coba panggil bisa gak", someName)
}
