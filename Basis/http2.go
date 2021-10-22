package main

import (
	"net/http"
	"fmt"
)

type dollars float32

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request){
	switch r.URL.Path{
		case "/list":
			for item, price := range db{
				fmt.Fprintf(w, "%s: %s\n", item,  price)
			}
		case "/price":
			item := r.URL.Query().Get("item")
			price, ok := db[item]
			if !ok{
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "not such item: %q\n", item)
				return
			}
			fmt.Fprintf(w, "%s\n", price)
			
		default:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "no such page: %s\n", r.URL)

	}
}