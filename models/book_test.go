package models

import (
	"fmt"
	"testing"
)

func TestGetById(t *testing.T){
	book := GetBookById("3d9aeb87-98fc-4008-b7fe-35f58a8a3b67")
	if book == nil {
		t.Errorf("error bookis empty");
	}
}

func TestAddBook(t *testing.T){
	tmp := Book{
		ISBF: "test",
		Title: "test title",
		Page: 24,
		Author: 1,
	}
	message, err := AddBook(tmp)

	if message == "" && err != nil {
		t.Errorf("%v", err)
	}
}

func TestUpdateBook(t *testing.T){
	book := GetBookById("3d9aeb87-98fc-4008-b7fe-35f58a8a3b67")

	if book == nil{
		t.Errorf("invalid book in the records")
	}

	book.Title = "test*"
	book.ISBF = "test*"
	book.Page = 101

	fmt.Println(book)
	if book, err := UpdateBook(*book); book == nil && err != nil{
		t.Errorf("%v",err)
	}
}

func TestDeleteById(t *testing.T){
	if msg, err := DeleteById("3d9aeb87-98fc-4008-b7fe-35f58a8a3b67"); msg == "" && err != nil{
		t.Errorf("error while deleting the book")
	}
}