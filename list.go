package main

import (
	"container/list"
	"fmt"
	"strings"
)

func main() {
	var data *list.List = list.New()
	data.PushBack("Andi")
	data.PushBack("Luiz")
	data.PushBack("Val")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// Manual Check

	fmt.Println("\n", strings.Repeat("#", 40), "\n")

	var head *list.Element = data.Front()
	fmt.Println(head.Value)

	next := head.Next()
	fmt.Println(next.Value)
	next = next.Next()
	fmt.Println(next.Value)
}

// 68
