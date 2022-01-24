package main

import (
	"OZON/internal/db"
	"OZON/internal/handlers"
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	"github.com/eknkc/basex"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	ALPHABET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "example"
	dbname   = "postgres"
)

func main() {
	dataFlag := flag.String("db", "im", "InMemory or Postgres database")
	flag.Parse()

	var data db.DB

	switch *dataFlag {
	case "im":
		data = db.NewInMemory()
		fmt.Println("Picked InMemory")
	case "psql":
		psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
		conn, err := sql.Open("postgres", psqlconn)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		data = db.NewPostgresDB(conn)
		fmt.Println("Picked PSQL")
	default:
		fmt.Println("Use im for InMemory or psql for Postgres")
		return
	}

	encoder, err := basex.NewEncoding(ALPHABET)
	if err != nil {
		fmt.Println("Cannot make encoder")
		return
	}

	add := handlers.NewAddHandler(data, encoder)
	get := handlers.NewGetHandler(data)

	r := mux.NewRouter()
	r.Handle("/", add).Methods("POST")
	r.Handle("/{key}", get).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
