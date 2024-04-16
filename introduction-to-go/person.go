package main

import (
	"fmt"
)

func main() {
	Person := map[string]string{"name": "John", "age": "40"}

	fmt.Println(Person["name"]) // accessing individual value from the map

	// iterating through the map
	for key, value := range Person {
		fmt.Printf("%q : %s\n", key, value)
	}
}
