package main

import "fmt"

func main() {
	firstname := "John"
	lastname := "Doe"

	// Println
	fmt.Println("Hello '", firstname, lastname, "'")

	// Printf
	fmt.Printf("Hello '%s %s'\n", firstname, lastname)
}
