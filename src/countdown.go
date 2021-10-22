package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	// fmt.Println("Connecting countdown.")
	// tick := time.Tick(1 * time.Second)
	// for countdown := 10; countdown > 0; countdown-- {
	// 	fmt.Println(countdown)
	// 	<-tick
	// }

	fmt.Println("Connecting countdown. Press return to abort. ")
	select {
	case <-time.After(10 * time.Second):
		// Do nothing
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}

	launch()
}

func launch() {
	fmt.Println("Lift off")
}
