package main

import (
	"fmt"
)

func main() {
	var item string
	var count int

	// Get user input
	fmt.Print("Enter the item you want to count: ")
	fmt.Scanln(&item)

	fmt.Print("How many times do you want to count it? ")
	fmt.Scanln(&count)

	// Display the count
	fmt.Printf("You have counted %s %d times.\n", item, count)
}
