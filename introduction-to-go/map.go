/*
Maps in go is created using the map keyword
map[key datatype][value data type]{}
*/
package main

import (
	"fmt"
)

func main() {
	studentDetails := map[string]string{
		"name":   "Jon",
		"age":    "20",
		"course": "Go backend course",
	}

	fmt.Println(studentDetails)

	// accessing key values from the map
	fmt.Println(studentDetails["name"])
}
