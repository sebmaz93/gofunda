package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":  4.99,
		"pie":   1.99,
		"salad": 3.99,
		"milk":  2.99,
	}

	fmt.Println(menu["soup"], menu)

	// looping over map
	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

	// Int as key in map
	phonebook := map[int]string{
		1: "John", 2: "Jane", 3: "Jack",
	}
	fmt.Println(phonebook[1])

	phonebook[4] = "Tom"
	fmt.Println(phonebook[4])
}
