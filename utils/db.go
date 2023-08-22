package utils

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

/*
using singleton design pattern for database connection
*/

var db *sql.DB

func GetInstance() *sql.DB{
	if db == nil{
		connStr := "postgres://aymane:aymane@123@localhost/book_go?sslmode=disable"
		db, err := sql.Open("postgres", connStr)
		if err != nil{
			log.Fatalf("%v", err)
		}
		return db
	}
	return db
}