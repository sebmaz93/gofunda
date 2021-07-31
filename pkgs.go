package main

import (
	"fmt"
	"sort"
)

func main() {
	// greetings := "hello seb"

	// fmt.Println(strings.Contains(greetings, "se"))
	// fmt.Println(strings.ReplaceAll(greetings, "seb", "jax")) // returns a new string
	// fmt.Println(strings.ToUpper(greetings))
	// fmt.Println(strings.Index(greetings, "ll"))
	// fmt.Println(strings.Split(greetings, " "))

	ages := []int{1, 22, 35, 4, 55}

	sort.Ints(ages) // this mutates the original slice
	fmt.Println(ages)

	index := sort.SearchInts(ages, 22)
	fmt.Println(index)

	names := []string{"seb", "jax", "joe", "jane"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "seb"))
}
