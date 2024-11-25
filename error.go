package main

import "errors"

var (
	ValidationError = errors.New("Validation error")
	NotFoundError   = errors.New("Not found")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "123" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := GetById("1")

	if err != nil {
		if errors.Is(err, ValidationError) {
			println("Validation error")
		} else if errors.Is(err, NotFoundError) {
			println("Not found")
		} else {
			println("Unknown error")
		}
	}

}
