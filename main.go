package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "hello there friends!"

	fmt.Println(strings.Contains(greeting, "hello"))         // will return true, like includes method in javascript, cool!
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi")) // will replace the second argument with the third argument
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll")) // find the index position of ll
	fmt.Println(strings.Split(greeting, " "))

	// the original value is unchanged
	fmt.Println("the original value is: ", greeting)

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(ages) // sort integer and the original value will changed
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30) // return position of 30 inside the ages
	fmt.Println(index)

	names := []string{"yoshi", "mario", "lalala"}
	sort.Strings(names) // will sort and changed the original value
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "mario")) // will find the position of mario
}
