package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Welcome to Nomadcoin\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("explorer:	Start the HTML Explorer\n")
	fmt.Printf("rest:		Start the REST API (recommended)\n\n")
	os.Exit(0)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}
}
