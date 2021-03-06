package handlers

import (
	"OZON/internal/db"
	"encoding/binary"
	"fmt"
	"github.com/eknkc/basex"
	"hash/fnv"
	"net/http"
)

type AddHandler struct {
	data    db.DB
	encoder *basex.Encoding
}

func NewAddHandler(data db.DB, encoder *basex.Encoding) *AddHandler {
	return &AddHandler{data: data, encoder: encoder}
}

func (h AddHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	url := r.Form.Get("url")
	if url == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	str := URLEncode(url, h)
	h.data.AddURL(url, str)
	w.Write([]byte("localhost:8080/" + str))
}

func URLEncode(url string, h AddHandler) string {
	hash := fnv.New64a()
	hash.Write([]byte(url))
	hashSum := hash.Sum64()
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, hashSum)
	str := h.encoder.Encode(b)
	str = str[:10]
	return str
}
