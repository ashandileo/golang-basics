package main

import "fmt"

func main() {
	myBill := newBill("ashandi's bill")

	myBill.addItem("soup", 4.50)
	myBill.addItem("coffe", 3.50)
	myBill.updateTip(10)

	fmt.Println(myBill.format())
}
