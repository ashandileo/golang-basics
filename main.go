package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Printf("Good morning %v \n", name)
}

func sayBye(name string) {
	fmt.Printf("Goodbye %v \n", name)
}

func cycleNames(names []string, callback func(string)) {
	for _, value := range names {
		callback(value)
	}
}

func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}

func main() {
	sayGreeting("Ashandi")
	sayGreeting("Lutpi")
	sayBye("Ashandi")

	cycleNames([]string{"caca", "cici", "cucu"}, sayGreeting)
	cycleNames([]string{"caca", "cici", "cucu"}, sayBye)

	area1 := circleArea(10.5)
	area2 := circleArea(15)
	fmt.Println(area1, area2)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", area1, area2)
}
