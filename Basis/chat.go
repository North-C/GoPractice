package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string // out channel

var (
	leaving  = make(chan client)
	entering = make(chan client)
	message  = make(chan string) // in channel
)

func broadcaster() {
	clients := make(map[client]bool)

	for {
		select {
		case msg := <-message:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		case cli := <-entering:
			clients[cli] = true
		}
	}
}

func handleConn(conn net.Conn) {
	// information out channel
	ch := make(chan string)
	go clientWriter(conn, ch)

	// inform the client is coming by entering channel
	who := conn.RemoteAddr().String()
	ch <- who + ": "
	message <- who + "has arrived"
	entering <- ch
	// read every msg from the client, broadcast to global channel
	input := bufio.NewScanner(conn)
	for input.Scan() {
		if input.Err() != nil {
			fmt.Fprintf(os.Stderr, "Read conn %v failed: %s", conn, input.Err())
		}

		message <- who + ":" + input.Text()
	}

	// leaving channel to inform the client is leaving
	leaving <- ch
	message <- who + " left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
