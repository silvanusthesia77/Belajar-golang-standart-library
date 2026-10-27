package main

import (
	"flag"
	"fmt"
)

func main() {
	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "root", "database password")
	var host *int = flag.Int("host", 0, "database host")
	var port *string = flag.String("port", "root", "database port")

	flag.Parse()
	fmt.Println("Username :", *username)
	fmt.Println("Password :", *password)
	fmt.Println("Post :", *host)
	fmt.Println("Port :", *port)

}

// 64
