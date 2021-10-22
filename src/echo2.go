package main

import (
	"fmt"
	"os"
)

func main() {
	s, seg := "", ""

	for _, arg := range os.Args[1:] {
		s += seg + arg
		seg = " "
	}
	fmt.Println(s)
}
