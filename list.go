package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()

	data.PushBack("Hello")
	data.PushBack("World")
	data.PushBack("Golang")

	var head *list.Element = data.Front()
	fmt.Println(head.Value) // Hello

	next := head.Next()
	fmt.Println(next.Value) // World

	next = next.Next()
	fmt.Println(next.Value) // Golang

	next = next.Prev()
	fmt.Println(next.Value) // World

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}
