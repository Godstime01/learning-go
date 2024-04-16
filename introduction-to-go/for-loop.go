package main

import "fmt"

func main() {
	// slice of string of sharks
	sharks := []string{"hammerhead", "great white", "dog fish", "frilled"}

	// for loop to iterates over sharks list
	for _, sharks := range sharks {
		fmt.Println(sharks)
	}
}
