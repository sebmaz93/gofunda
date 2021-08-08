package main

import "fmt"

func updateVar(x *string) {
	*x = "seb"
}

func main() {
	name := "astgh"

	// updateVar(name)

	fmt.Println("name : ", name)

	nameP := &name
	fmt.Println("memory addr", nameP)
	fmt.Println("value of this mem addr", *nameP)

	updateVar(nameP)

	fmt.Println("name : ", name)
}
