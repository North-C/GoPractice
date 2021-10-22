package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch8/thumbnail"
)

func makeThumbnails(filenames []string){
	for _, f := range filenames{
		if _, err := thumbnail.ImageFile(f); err!= nil{
			log.Println(err)
		}
	}
}

func makeThumbnails2(filenames []string){
	for _, f:= range filenames{
		go thumbnail.ImageFile(f)		// Not Correct
	}
}

func makeThumbnails3(filenames []string){
	ch := make(chan struct{})

	for _,f := range filenames{
		go func(f string){
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(f)
	}
}

func makeThumbnails4(filenames []string) error{
	errors := make(chan error)

	for _, f := range filenames{
		go func(f string){
			_, err := thumbnail.ImageFile(f) 
			errors <- err
		}(f)
	}


	for range filenames{
		if err := <- errors; err != nil{
			return err
		}
	}
	return nil
}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error){
	type item struct{
		thumbfile string 
		err error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames{
		go func(f string){
			var it, item
			it.thumbfile, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames{
		it := <- ch
		if it.err != nil{
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}

func thumbnail6(filenames <-chan string)(filename []string, err error){
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames{
		wg.Add(1)

		//worker
		go func(f string){
			defer wg.Done()
			thumb, err := thumbnail.ImageFile(f)
			if err != nil{
				log.Println(err)
				return
			}
			info, _ := os.Stat(thumb)
			size <- info.Size()
		}(f)
	}

	go func(){
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size  := range sizes{
		total += size
	}
	return total

}
