package main

import (
	"fmt"
	"net"
	"os"
)



import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form{
		fmt.Println("key", k)
		fmt.Println("value", v)
	}
	fmt.Fprintf(w, "Hello New man")
}

func login(w http.ResponseWriter, r *http.Request){
		fmt.Println("method:", r.Method)
		if r.Method == "GET"{
			t, _ := template.ParseFiles("login.gtpl")
			log.Println(t.Execute(w, nil))
		}else{
			fmt.Println("username:", r.Form[username])
			fmt.Println("password:", r.Form[password])
		}
}

func main(){
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err!=nil{
		log.Fatal("ListenAndServe: ", err)
	}
}
