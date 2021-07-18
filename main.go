package main

import "fmt"

var score = 99.5

func main() {

	// Map -> allows us to store key value pairs where the key and values can be different type
	// But in a single map all of the keys and value must be the same type

	menu := map[string]float64{ // type of key and type of value
		"soup":          4.99,
		"pie":           7.99,
		"salad":         6.99,
		"toffe pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping map
	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

	// int as key type
	phoneBook := map[int]string{
		8811882850: "ashandi",
		8828128910: "alfalah",
		9219179312: "dani",
	}

	fmt.Println(phoneBook)
	fmt.Println(phoneBook[8811882850])

	// Update map value
	phoneBook[8811882850] = "leonadi"
	fmt.Println(phoneBook)
}
