package main

// package for formatting strings and priting message to console
import "fmt"

// automatically will be fired by go because of naming main()
func main() {
	fmt.Println("Hello, Seb")

	// strings
	var dog string = "jax"
	var cat = "tom"
	var wife string

	wife = "astgh"

	dog2 := "zeus" // this syntax only works inside func

	fmt.Println(dog, cat, wife, dog2)

	// ints
	var age int = 20
	var num1 int8 = 100

	// bits and momory
	// int8 (-128 to 127)
	// int16 (-32768 to 32767)
	// int32 (-2147483648 to 2147483647)
	// int64 (-9223372036854775808 to 9223372036854775807)
	// int (same as int32)
	// uint8 (0 to 255) positive only

	// floats
	var float1 float32 = 1.5

	fmt.Println(age, num1, float1)

	// arrays and slices
	var arr1 [4]int = [4]int{1, 2, 3, 4} // array of 4 ints arrays are fixed size
	arr1[2] = 22

	fmt.Println(arr1, "length of arr", len(arr1))

	// slices use array under the hood
	var scores = []int{100, 200, 300, 400}
	scores = append(scores, 500) // appends to end of slice) and returns new slice

	fmt.Println(scores)

	// slice ranges
	rangeOne := scores[1:3]  // from 1 to 3
	rangeTwo := scores[:3]   // from beginning to 3rd element
	rangeThree := scores[1:] // from 1st element to end
	fmt.Println(rangeOne, rangeTwo, rangeThree)
}
