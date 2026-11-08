package main

import (
	"fmt"
	"time"
)

func main() {
	var duration time.Duration = time.Second * 100
	var duration1 time.Duration = time.Millisecond * 10
	var duration2 time.Duration = duration + duration1

	fmt.Println(duration)
	fmt.Println(duration1)
	fmt.Printf("Durasi : %d", duration2)
}

//  71
