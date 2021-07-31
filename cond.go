package main

import "fmt"

func main() {
	age := 45

	if age < 30 {
		fmt.Println("you are young")
	} else if age > 30 {
		fmt.Println("You are old enough to drive")
	} else {
		fmt.Println("You are exactyl 30")
	}

	names := []string{"John", "Paul", "George", "Ringo", "Seb"}

	for index, value := range names {
		if index == 2 {
			fmt.Println(index, value)
			continue // continue the loop and not fire what is after the continue
		}

		if index > 2 {
			fmt.Println("Breaking at pos", index)
			break // break the loop and exit the loop
		}

		fmt.Printf("the value at pos %v is %v \n", index, value)
	}
}
