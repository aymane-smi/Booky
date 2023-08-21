package models

import (
	"fmt"

	"github.com/aymane-smi/api-test/utils"
	"github.com/google/uuid"
)

//can be used as DTO
type Book struct{
	Id string 
	ISBF string 
	Title string 
	Page int 
	Author int
}

//get book using his ID

func GetBookById(id string) *Book{
	db := utils.GetInstance()
	stmt, err := db.Prepare("SELECT * FROM books WHERE id = $1")
	if err != nil{
		utils.Log.Error(err.Error())
		return nil
	}

	row := stmt.QueryRow(id)
	var tmp_book Book
	if err := row.Scan(&tmp_book.Id, &tmp_book.ISBF, &tmp_book.Title, &tmp_book.Page, &tmp_book.Author); err != nil{
		utils.Log.Error(err.Error())
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
		utils.Log.Error(err.Error())
		return "", err
	}
	if _, err := stmt.Exec(newUUID.String(), b.ISBF, b.Title, b.Page, b.Author); err != nil{
		utils.Log.Error(err.Error())
		return "", err
	}
	return "new row inserted in books", nil

}

//add a book by passing a book to the function for testing purpose

func AddBookTest(b Book) (string, error){
	newUUID := "test_id"
	db := utils.GetInstance()
	stmt, err := db.Prepare("INSERT INTO books(id, isbf, title, page, author_id) VALUES($1, $2, $3, $4, $5)")
	if err != nil{
		utils.Log.Error(err.Error())
		return "", err
	}
	if _, err := stmt.Exec(newUUID, b.ISBF, b.Title, b.Page, b.Author); err != nil{
		utils.Log.Error(err.Error())
		return "", err
	}
	return "new row inserted in books", nil

}

//update a book by passing new tems to change in the records

func UpdateBook(b Book) (*Book, error){
	db := utils.GetInstance()
	stmt, err := db.Prepare("UPDATE books SET title = $1, isbf = $2, page = $3 WHERE id = $4")
	if err != nil {
		utils.Log.Error(err.Error())
		return nil, err
	}

	if _, err := stmt.Exec(b.Title, b.ISBF, b.Page, b.Id); err != nil{
		utils.Log.Error(err.Error())
		return nil, err
	}

	return &b, err
}

func DeleteById(id string) (string, error){
	db := utils.GetInstance()
	stmt, err := db.Prepare("DELETE FROM books WHERE id = $1")
	if err != nil{
		utils.Log.Error(err.Error())
		return "", err
	}
	if _, err := stmt.Exec(id); err != nil{
		utils.Log.Error(err.Error())
		return "", err
	}

	return fmt.Sprintf("book with id '%s' was deleted", id), nil
}