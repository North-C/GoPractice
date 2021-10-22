package main

import (
	"fmt"
	"bytes"
)

type IntSet struct{
	words []uint64
}

func main(){

}

func (s *IntSet)Has(x int) bool{
	word, bit := x/64, uint(x%64)

	if word < len(s.words) && s.words[word]&(1<<bit) !=0{
		return true
	}
	return false

	//simple way to return:
	//return  word < len(s.words) && s.words[word]&(1<<bit) !=0
}

func (s *IntSet) Len() int{
	return	len(s.words)
}

func (s *IntSet)Remove(x int){
	word, bit := x/64, x%64
	if word > len(s.words){
		return
	}
	s.words[word] &^= 1<<bit   
	//s.words[word]  = uint64(bits+1) ^ (s.words[word])
}


func (s *IntSet) Clear(){
	s.words = nil
}

func (s *IntSet)Copy() *IntSet{

	p := &IntSet{}
	copy(p.words, s.words)
	return p
}

func (s *IntSet) String() string{
	var buf bytes.Buffer
	buf.WriteByte('{') 				//Buffer need to close ???
	for i, word := range s.words{
		if word == 0{
			continue
		}

		for j := 0; j <64; j++{
			if word&(1<<uint(j)) != 0{
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}