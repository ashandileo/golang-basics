package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// func to create new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// format the bill
func (b *bill) format() string { // just add reference to parameter and golang will automatically set the update action as reference not a copy
	fs := "Bill breakdown \n"
	var total float64 = 0

	// list items
	for key, value := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", key+":", value)
		total += value
	}

	// add tip
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "tip: ", b.tip)

	// total
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "total: ", total+b.tip)
	return fs
}

// update the tip
func (b *bill) updateTip(tip float64) { // (*bill) reference to bill -> pointer to update the original value
	b.tip = tip
}

// add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("bill was saved to file")
}
