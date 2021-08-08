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

// make new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// format the bill
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-15v ...$%v \n", k+":", v)
		total += v
	}

	// tip
	fs += fmt.Sprintf("%-15v ...$%v \n", "tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-15v ...$%0.2f \n", "total:", total+b.tip)

	return fs
}

// update the tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add an item to the bill
func (b *bill) addItem(item string, price float64) {
	b.items[item] = price
}

// save bill to file
func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		fmt.Println("Error: saving the file")
		panic(err)
	}
	fmt.Println("Bill saved")
}
