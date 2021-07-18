package main

import "fmt"

func updateMenu(menu map[string]float64) {
	menu["coffe"] = 2.99
}

func main() {

	// Group B => slices, maps, functions
	// will change the original value
	menu := map[string]float64{
		"pie":       5.95,
		"ice cream": 3.99,
	}

	updateMenu(menu)
	fmt.Println(menu) // we will see coffe as the new value
}
