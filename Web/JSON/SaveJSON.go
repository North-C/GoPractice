package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/check.v1"
)


type Person struct{
	Name Name
	Email []Email
}

type Name struct{
	Family string
	Personal string
}

type Email struct{
	Kind string
	Address string
}

func main(){
	person := Person{
		Name: Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{Email{Kind:"home", Address: "jan@newmarch.name"},
			Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}

	saveJSON("person.json", person)
}


func saveJSON(fileName string, key interface{}){
	outFile, err := os.Create(fileName)
	checkError(err)
	encoder := json.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
	outFile.Close()
}

func loadJSON(fileName string, key interface{}){
	inFile, err := os.Open(fileName)
	checkError(err)
	decoder := json.NewDecoder(inFile)
	err = decoder.Decode(inFile)
	checkError(err)
	inFile.Close()
}

func checkError(err error){
	if err != nil{
		fmt.Println("Fatal error", err.Error())
		os.Exit(1)
	}
}

