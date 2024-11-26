package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"100"`
}

type Person struct {
	Name  string `required:"true" max:"100"`
	Age   string `required:"true" max:"100"`
	Email string `required:"true"`
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("Type: ", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		var structField reflect.StructField = valueType.Field(i)

		fmt.Println(structField.Name, "with type : ", structField.Type)
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))
	}

}

func IsValid(value any) (result bool) {
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""

			if result == false {
				return result
			}
		}
	}

	return true
}

func main() {
	readField(Sample{"Hello"})
	readField(Person{"Akmal", "", ""})

	person := Person{
		Name: "Akmal",
		Age:  "20",
		// Email: "test@gmail.com",
	}

	fmt.Println(IsValid(person))
}
