package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Eko Kurniawan", "Eko"))
	fmt.Println(strings.Split("Eko Kurniawan", " "))
	fmt.Println(strings.ToLower("Eko Kurniawan Solossa"))
	fmt.Println(strings.ToUpper("Eko Kurniawan Solossa"))
	fmt.Println(strings.Trim("    Eko Kurniawan Solossa   ", " "))
	fmt.Println(strings.ReplaceAll("Eko Kurniawan Eko Solossa Eko Kambu", "Eko", "Thobiaz"))

}

// 65
