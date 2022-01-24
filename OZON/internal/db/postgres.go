package db

import (
	"database/sql"
	"fmt"
)

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresDB(data *sql.DB) *PostgresDB {
	return &PostgresDB{db: data}
}

func (p PostgresDB) AddURL(url, key string) error {
	insertStmt := `INSERT INTO "urls"("key", "url") VALUES($1, $2)`
	_, err := p.db.Exec(insertStmt, key, url)
	return err
}

func (p PostgresDB) GetURL(key string) (string, error) {

	var url string
	selectStmt := `SELECT "url" FROM "urls" WHERE "key"=$1`
	row := p.db.QueryRow(selectStmt, key)
	err := row.Scan(&url)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return url, err
	case nil:
		return url, err
	default:
		fmt.Println("Error")
		return url, err
	}
}
