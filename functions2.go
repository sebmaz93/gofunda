package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	return n[0:1], n[1:2]
}

func getUpperCase(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string

	for _, value := range names {
		initials = append(initials, strings.ToUpper(value[:1]))
	}
}

func main() {
	fmt.Println(getInitials("Gopher"))
}
