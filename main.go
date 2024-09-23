package main

import (
	"cache-proxy/proxy"
	"flag"
	"log"
)

func main() {
	origin := flag.String("origin", "", "origin ")
	flag.Parse()

	if *origin == "" {
		panic("origin must containe url address")
	}

	if err := proxy.HttpCall(*origin); err != nil {
		log.Fatal(err.Error())
	}

}
