package main

import (
	"fmt"
	"time"
)

func main() {
	var duration1 time.Duration = 150 * time.Millisecond
	var duration2 time.Duration = 10 * time.Second
	var duration3 time.Duration = 1 * time.Minute

	fmt.Println("Duration 1 is", duration1)
	fmt.Println("Duration 2 is", duration2)
	fmt.Println("Duration 3 is", duration3)
}
