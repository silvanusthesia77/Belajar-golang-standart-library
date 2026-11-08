package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Students struct {
	Name    string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
}
type Teacher struct {
	Name    string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
}

func readFiel(value any) {
	reflectValue := reflect.TypeOf(value)
	for i := 0; i < reflectValue.NumField(); i++ {
		vallu := reflectValue.Field(i)
		fmt.Println(vallu.Name, "type", vallu.Type)
	}
}
func main() {
	readFiel(Students{"Wanus", "", ""})
	readFiel(Teacher{"thoby", "", ""})
	fmt.Println("\n", strings.Repeat("#", 30), "\n")
	var anggota = Students{
		Name:    "Toby",
		Email:   "Biba@gmail.com",
		Address: "Jln. Jogja",
	}
	fmt.Println(isValit(anggota))
	var teachers = Students{
		Name:    "Toby",
		Email:   "Biba@gmail.com",
		Address: "Jln. Jogja",
	}
	fmt.Println(isValit(teachers))
}
func isValit(value any) bool {
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			return data != ""
		}
	}
	return true
}
