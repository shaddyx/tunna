package main

import (
	"flag"
	"fmt"
)

func main() {
	serverMode := flag.Bool("server", false, "Run software in a server mode")
	flag.Parse()
	if *serverMode {
		fmt.Println("Server mode")
	} else {
		err := InitClient()
		if err != nil {
			fmt.Println(err)
		}
	}
}
