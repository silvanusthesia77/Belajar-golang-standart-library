package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "thoby, junior, parviddey\n" + "luiz, arthur, kndmy\n" + "reza, gledy, rawar"
	read := csv.NewReader(strings.NewReader(csvString))
	for {
		recorded, err := read.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(recorded)
	}
}
