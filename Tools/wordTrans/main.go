package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "Go language", "help info")
	flag.StringVar(&name, "n", "Go language", "help info")
	flag.Parse()

	log.Printf("name: %s", name)
}
