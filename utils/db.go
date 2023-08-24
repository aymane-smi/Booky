package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

/*
using singleton design pattern for database connection
*/

var db *sql.DB

func GetInstance() *sql.DB{

	if db == nil{
		connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
		db, err := sql.Open("postgres", connStr)
		if err != nil{
			log.Fatalf("%v", err)
		}
		return db
	}
	return db
}