package main

import "fmt"

func updateName(name *string) {
	*name = "leonadi"
}

func main() {
	// To update original value in group A
	// you can pass the pointers as argument

	name := "ashandi"

	fmt.Println(name) //ashandi

	updateName(&name)

	fmt.Println(name) //leonadi
}
