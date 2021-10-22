package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	h    bool
	v, V bool
	t, T bool
	q    *bool
	s    string
	p    string
	c    string
	g    string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&v, "v", false, "show version and exit")
	q = flag.Bool("q", false, "suppress non-error messages during configurations testing.")

	flag.StringVar(&s, "s", "", "send `signal` to a master process")
	flag.StringVar(&c, "c", "/usr/local/nginx", "set `prefix` path")

	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0
	Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]
	
	Options:`)

	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
	}
}
