package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/studokim/resolver/internal"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		if args[1] == "--listen" || args[1] == "-l" {
			port, err := strconv.Atoi(args[2])
			internal.HandleFatal(err)
			internal.Listen(port)
		} else {
			var resolver internal.Resolver
			resolver.Init()
			address := resolver.Resolve(args[1])
			fmt.Println(address)
		}
	} else {
		fmt.Println("Usage: ./resolver <example.com>")
		fmt.Println("or     ./resolver --listen <port>")
		fmt.Println("or     ./resolver -l       <port>")
	}
}
