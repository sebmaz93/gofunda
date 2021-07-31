package main

import "fmt"

func main() {
	// while loop
	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x : ", x)
	// 	x++
	// }

	// for loop
	// for i :=0; i < 5; i++ {
	// 	fmt.Println("value of i : ", i)
	// }

	names := []string{"John", "Paul", "George", "Ringo"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, value := range names {
	// 	fmt.Printf("this position at %v is %v \n", index, value)
	// }

	for _, value := range names {
		fmt.Printf("this value is %v \n", value)
	}

}
