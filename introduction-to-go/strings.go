package main

import (
	"fmt"
	"strconv"
)

func main() {
	var firstName string = "Kingsley"
	var lastName string = "Londonway"

	var number int = 40
	// convert number to string
	fmt.Println("Convert number to string " + strconv.Itoa(number))

	// string concatenation
	fmt.Println(firstName + lastName)
}
