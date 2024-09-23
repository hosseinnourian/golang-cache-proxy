package main

import (
	"flag"
	"fmt"
)

func main() {
	origin := flag.String("origin", "", "origin ")
	flag.Parse()

	if *origin == "" {
		panic("origin must containe url address")
	}

	fmt.Println(*origin)

}
