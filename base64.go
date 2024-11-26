package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "Hello World"

	// Encode
	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println("Encoded:", encoded)

	// Decode
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Decoded:", string(decoded))
	}
}
