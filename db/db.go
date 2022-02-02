package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {
	connStr := "user=postgres dbname=alura_loja password=.A9l13ed0m7., host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
