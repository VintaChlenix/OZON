package main

import (
	"OZON/internal/db"
	"OZON/internal/handlers"
	"fmt"
	"github.com/eknkc/basex"
	"github.com/gorilla/mux"
	"net/http"
)

const ALPHABET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

func main(){
	encoder,err:=basex.NewEncoding(ALPHABET)
	if err!=nil{
		fmt.Println("Cannot make encoder")
		return
	}
	data:= db.NewInMemory()
	add := handlers.NewAddHandler(data, encoder)
	get := handlers.NewGetHandler(data)
	r := mux.NewRouter()
	r.Handle("/", add).Methods("POST")
	r.Handle("/{key}", get).Methods("GET")
	http.Handle("/",r)
	http.ListenAndServe(":8080",nil)
}
