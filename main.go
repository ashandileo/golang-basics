package main

import (
	"fmt"
	"strings"
)

func getInitials(name string) (string, string) {
	uppercaseName := strings.ToUpper(name)
	names := strings.Split(uppercaseName, " ")

	var initials []string
	for _, value := range names {
		initials = append(initials, value[:1]) //get the first letter [:1]
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	firstName1, secondName1 := getInitials("ashandi leo")
	fmt.Println(firstName1, secondName1)

	firstName2, secondName2 := getInitials("test jared")
	fmt.Println(firstName2, secondName2)

	firstName3, secondName3 := getInitials("richard")
	fmt.Println(firstName3, secondName3)
}
