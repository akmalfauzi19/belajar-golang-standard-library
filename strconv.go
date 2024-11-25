package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")

	if err != nil {
		fmt.Println("errror", err.Error())
	} else {
		fmt.Println("res : ", boolean)
	}

	// Atoi
	number, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("errror", err.Error())
	} else {
		fmt.Println("res : ", number)
	}

	// srtconv format
	binary := strconv.FormatInt(1000, 2)
	fmt.Println("binary : ", binary)

	var stringInt string = strconv.Itoa(1000)
	fmt.Println("stringInt : ", stringInt)
}
