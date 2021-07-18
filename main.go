package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch option {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		floatPrice, err := strconv.ParseFloat(price, 64) //float64
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}

		b.addItem(name, floatPrice)

		fmt.Println("Item added - ", name, floatPrice)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		floatTip, err := strconv.ParseFloat(tip, 64) //float64
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}

		b.updateTip(floatTip)

		fmt.Println("Tip added - ", floatTip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("you saved the file - ", b.name)
	default:
		fmt.Println("that was not a valid option...")
		promptOptions(b)
	}
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin) // get user input with standart input meaning the terminal

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
}
