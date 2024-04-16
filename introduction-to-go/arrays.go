/*
GO arrays size are static and is allocated before compilation time
*/

package main

import (
	"fmt"
)

func main() {
	names := [3]string{"Jon", "Mikel", "Love"}
	fmt.Println(names)
}
