package main

import (
	"database/sql"
	"net/http"
)

type MyContext struct {
	db *sql.DB
}

func (m *MyContext)handle(w http.ResponseWriter, r*http.Request){

}

func main(){
	myctx := NewMyContext()
	http.HandleFunc("/",myctx.handle)
}