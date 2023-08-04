package models

import (
	"fmt"
	"github.com/aymane-smi/api-test/utils"
	"github.com/google/uuid"
)

type Book struct{
	id string
	ISBF string
	title string
	page int
	author int
}

//get book using his ID

func GetBookById(id string) *Book{
	db := utils.GetInstance()
	stmt, err := db.Prepare("SELECT * FROM books WHERE id = $1")
	if err != nil{
		fmt.Println(err)
		return nil
	}

	row := stmt.QueryRow(id)
	var tmp_book Book
	if err := row.Scan(&tmp_book.id, &tmp_book.ISBF, &tmp_book.title, &tmp_book.page, &tmp_book.author); err != nil{
		fmt.Println(err)
		return nil
	}
	return &tmp_book
	
}

//add a book by passing a book to the function

func AddBook(b Book) (string, error){
	newUUID := uuid.New()
	db := utils.GetInstance()
	stmt, err := db.Prepare("INSERT INTO books(id, isbf, title, page, author_id) VALUES($1, $2, $3, $4, $5)")
	if err != nil{
		panic(err)
		return "", err
	}
	if _, err := stmt.Exec(newUUID.String(), b.ISBF, b.title, b.page, b.author); err != nil{
		panic(err)
		return "", err
	}
	return "new row inserted in books", nil

}

//update a book by passing new tems to change in the records

func updateBook(b Book) (*Book, error){
	db := utils.GetInstance()
	stmt, err := db.Prepare("UPDATE books SET title = $1, isbf = $2, page = $3 WHERE id = $4")
	if err != nil {
		panic(err)
		return nil, err
	}

	if _, err := stmt.Exec(b.title, b.ISBF, b.page, b.id); err != nil{
		return nil, err
	}

	return &b, err
}

func DeleteById(id string) (string, error){
	db := utils.GetInstance()
	stmt, err := db.Prepare("DELETE FROM books WHERE id = $1")
	if err != nil{
		panic(err)
		return "", err
	}
	if _, err := stmt.Exec(id); err != nil{
		panic(err)
		return "", err
	}

	return fmt.Sprintf("book with id '%s' was deleted", id), nil
}