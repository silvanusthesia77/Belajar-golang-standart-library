package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}
type Person struct {
	Name   string `required:"true" max:"10"`
	Email  string `required:"true" max:"10"`
	Adress string `required:"true" max:"10"`
}

func readField(value any) {
	var readfield reflect.Type = reflect.TypeOf(value)
	fmt.Println("Reflect Type", readfield.Name())
	for i := 0; i < readfield.NumField(); i++ {
		var structfield reflect.StructField = readfield.Field(i)
		fmt.Println(structfield.Name, " with type ", structfield.Type)
		fmt.Println(structfield.Tag.Get("required"))
		fmt.Println(structfield.Tag.Get("max"))
	}
}
func isValid(data interface{}) (resault bool) {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			data := reflect.ValueOf(data).Field(i).Interface()
			resault = data != ""
			if resault == false {
				return resault
			}
		}

	}
	return true
}
func main() {
	readField(Sample{"Biba"})
	readField(Person{"Junior", "", ""})

	person := Person{
		Name:   "Wanus",
		Email:  "Sil@gmail.com",
		Adress: "",
	}
	fmt.Println(isValid(person))
}
