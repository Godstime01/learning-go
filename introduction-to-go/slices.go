/*
GO slice size are flexible and can change in runtime
*/
package main

import "fmt"

func main() {
	names := []string{"Jon", "Mikel", "SlickKing"}
	fmt.Println(names)

	// to add more elements to the slice use the append(sliceVariable, 'item')
	names = append(names, "King")
	fmt.Println(names)
}
