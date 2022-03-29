package main

import "fmt"

func main() {
	var input int
	for {
		fmt.Println("Infinite Loop")

		fmt.Println("Give me a number?")
		fmt.Scanln(&input)
		if input%2 == 0 {
			fmt.Println("Odd")
		}
	}
}
