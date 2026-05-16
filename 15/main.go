package main

import "net/http"

func main() {
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		switch r.Method{
		case http.MethodGet:
		}

	})
}