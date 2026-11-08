package main

import (
	"fmt"
	"time"
)

func main() {
	var hour time.Time = time.Now()
	fmt.Println(hour.Local())

	utc := time.Date(2009, time.August, 5, 4, 6, 5, 3, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05:00"
	value := "2020-10-10 12:05:10:00"
	valuetime, err := time.Parse(formatter, value)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valuetime)
	}
}
