package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"hafizh", "akmal", "fauzi"}
	values := []int{1, 2, 3, 4, 5}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Contains(names, "hafizh"))
	fmt.Println(slices.Contains(values, 3))
	fmt.Println(slices.Index(names, "akmal"))
	fmt.Println(slices.Index(values, 4))
}
