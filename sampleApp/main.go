package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Options:")
	fmt.Println("1. Add item")
	fmt.Println("2. Add tips")
	fmt.Println("3. Save bill")

	option, _ := getInput("Choose an option: ", reader)

	switch option {
	case "1":
		name, _ := getInput("item name: ", reader)
		price, _ := getInput("item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("Error: the Price must a be number", err)
			promptOptions(b)
		}

		b.addItem(name, p)
		fmt.Println("item added - ", name, price)
		promptOptions(b)

	case "2":
		tip, _ := getInput("tip amount: ", reader)
		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("Error: the Tip must be a number", err)
			promptOptions(b)
		}

		b.updateTip(t)
		fmt.Println("tip updated - ", t)
		promptOptions(b)

	case "3":
		b.save()

	default:
		fmt.Println("Invalid option")
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}
