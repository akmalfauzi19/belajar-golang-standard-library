package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println("The time is", now)

	var utc time.Time = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Println("The time is", utc)
	fmt.Println("The month is", utc.Month())
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05" // 2006-01-02 15:04:05 is the reference time
	value := "2020-01-01 00:00:00"

	valTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error parsing time", err)
		return
	} else {
		fmt.Println("The time is", valTime)
	}
}
