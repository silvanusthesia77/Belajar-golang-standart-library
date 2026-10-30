package main

import (
	"fmt"
	"strconv"
)

func main() {
	Bool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(Bool)
	}

	resault, err := strconv.Atoi("9999")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(resault)
	}

	num := 42

	str := strconv.Itoa(num)
	fmt.Println(str)

	binarry := strconv.FormatInt(99999, 22)
	fmt.Println(binarry)

	var Strings string = strconv.Itoa(99999)
	fmt.Println(Strings)

	itoa := strconv.Itoa(88888)
	fmt.Println(itoa)
}
