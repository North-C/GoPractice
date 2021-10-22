package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		log.Println("done")
		done <- struct{}{}
	}()
	writedata(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine
}

func writedata(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
