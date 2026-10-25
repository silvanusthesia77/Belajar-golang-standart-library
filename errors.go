package main

import (
	"errors"
	"fmt"
)

var (
	validationError = errors.New("Validation Error")
	notFoundError   = errors.New("Not Found Error")
)

func GetById(id string) error {
	if id == "" {
		return validationError
	} else if id == "Biba" {
		return notFoundError
	} else {
		return nil
	}
}

func main() {
	err := GetById("Junior")
	if err != nil {
		if errors.Is(err, validationError) {
			fmt.Println("Validation Error Guys")
		} else if errors.Is(err, notFoundError) {
			fmt.Println("Not Found Errors")
		} else {
			fmt.Println("Unkwon Error")
		}
	}
}

// 62
