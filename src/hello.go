package main

import "fmt"

func main(){
	var x uint8 = 8
	fmt.Printf("%b\n",x)
	x &= (1<<2)
	fmt.Printf("%b",x)
}
