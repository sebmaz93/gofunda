package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Println("Hello", name)
}

func cycleNames(n []string, f func(string)) {
	for _, name := range n {
		f(name)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

var names = []string{"Bob", "Alice", "Carol"}

func main() {
	sayGreeting("Seb")

	cycleNames(names, sayGreeting)

	a1 := circleArea(10.5)

	fmt.Println("AREA 1", a1)

	fmt.Printf("AREA 1 with 3 decimal only = %0.3f", a1)
}
