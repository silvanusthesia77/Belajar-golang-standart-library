package main

import (
	"fmt"
	"os"
)

func main() {
	argmt := os.Args
	for _, v := range argmt {
		fmt.Println(v)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Hasil Error ", err.Error())
	}
}

//  63 Flag
