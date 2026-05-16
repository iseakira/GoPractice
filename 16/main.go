package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/",func(w http.ResponseWriter,r *http.Request){
		f,err := os.Open("/path/to/content.txt")
		if err != nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
		defer f.Close()
		io.Copy(w,f)
	})
}