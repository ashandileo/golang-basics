package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// func to create new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"pie": 5.99, "cake": 3.9},
		tip:   0,
	}

	return b
}

// format the bill
func (b bill) format() string {
	fs := "Bill breakdown \n"
	var total float64 = 0

	// list items
	for key, value := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", key+":", value)
		total += value
	}

	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total: ", total)
	return fs
}
