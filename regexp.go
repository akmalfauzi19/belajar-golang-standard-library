package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("akm([a-z])l")

	fmt.Println(regex.MatchString("akmal"))
	fmt.Println(regex.MatchString("akmzl"))
	fmt.Println(regex.MatchString("akmBl"))

	search := regex.FindAllString("akmal akmzl akmBl", -1)
	fmt.Println(search)
}
