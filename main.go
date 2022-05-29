package main

import (
	"fmt"
	"os"

	"github.com/studokim/resolver/internal"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		domain := args[1]
		address := internal.Resolve(domain)
		fmt.Println(address)
	} else {
		fmt.Println("Usage: ./resolver <example.com>")
	}
}
