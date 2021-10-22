package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func walkDir(dir string, n *sync.WaitGroup, filesizes chan int64) {
	defer n.Done()
	if cancelled() {
		return
	}

	for _, item := range dirents(dir) {
		if item.IsDir() {
			n.Add(1)
			subpath := filepath.Join(dir, item.Name())
			go walkDir(subpath, n, filesizes)
		} else {
			filesizes <- item.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}: // acquire token

	case <-done: //cancelled
		return nil
	}
	defer func() { <-sema }() //release a token

	f, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	defer f.Close()

	entry, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "readDir failed: %v\n", err)
		// return nil //Don't return: Readdir may return partial results.
	}
	return entry
}

func printPercentage(nfiles, nbytes int64) {
	fmt.Printf(" %d files: %.1f GB\n ", nfiles, float64(nbytes)/1e9)
}

var done = make(chan struct{})

//
func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {

	//
	roots := os.Args[1:]
	if len(roots) == 0 {
		fmt.Println("No files input. Default value is '.' ")
		roots = []string{"."}
	}

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	// Traverse the root
	filesizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, filesizes)
	}
	go func() {
		n.Wait()
		close(filesizes)
	}()

	// Print the results periodically

	tick := time.Tick(500 * time.Millisecond)
	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			//Drain fileSizes to allow existing goroutines to finish.
			for range filesizes {
				// do nothing
			}
			return
		case size, ok := <-filesizes:
			if !ok {
				break loop // whether filesizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printPercentage(nfiles, nbytes)
		}
	}
	printPercentage(nfiles, nbytes)
}
