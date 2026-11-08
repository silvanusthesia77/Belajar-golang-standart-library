package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regazz *regexp.Regexp = regexp.MustCompile(`e([a-z])o`)
	fmt.Println(regazz.MatchString("eko"))
	fmt.Println(regazz.MatchString("eKo"))
	fmt.Println(regazz.MatchString("elo"))
	fmt.Println(regazz.FindAllString("eko, elo, lukas , edo, matias, Eko, e1o, eno", 10))

}

// 73
