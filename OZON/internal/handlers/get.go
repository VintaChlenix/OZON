package handlers

import (
	"OZON/internal/db"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type GetHandler struct{
	data db.DB
}

func NewGetHandler(data db.DB) *GetHandler {
	return &GetHandler{data: data}
}

func (h GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key,ok:=vars["key"]
	if !ok{
		fmt.Println("Key Undefined")
		return
	}
	url, err := h.data.GetURL(key)
	if err!=nil{
		fmt.Println("Cannot get URL")
		return
	}
	w.Write([]byte(url))
}

