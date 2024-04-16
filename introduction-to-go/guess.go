package main

import "fmt"

const favColor string = "blue"

func main() {
	var guess string

	for {

		// ask user for guess
		fmt.Println("Guess my favourite color: ")

		//
		if _, err := fmt.Scanln(&guess); err != nil {
			fmt.Printf("%s\n", err)
			return
		}

		// Did they make the correct guess
		if favColor == guess {
			fmt.Printf("%s is my favourite color!\n", favColor)
			return
		}

		fmt.Printf("Sorry, %q is not my favourite color. guess again\n", guess)
	}
}
