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
		f := make(internal.Filter)
		f.Readconfig()
		c := make(internal.Cache)
		address := internal.Resolve(domain, &c, &f)
		fmt.Println(address)
	} else {
		fmt.Println("Usage: ./resolver <example.com>")
	}
}
