package models

import (
	"testing"
	"fmt"
)

func TestGetById(t *testing.T){
	book := GetById("3d9aeb87-98fc-4008-b7fe-35f58a8a3b67")
	if book == nil {
		t.Errorf("error bookis empty");
	}
}

func TestAddBook(t *testing.T){
	tmp := Book{
		ISBF: "test",
		title: "test title",
		page: 24,
		author: 1,
	}
	message, err := AddBook(tmp)

	if message == "" && err != nil {
		t.Errorf("%v", err)
	}
}

func TestUpdateBook(t *testing.T){
	book := GetById("3d9aeb87-98fc-4008-b7fe-35f58a8a3b67")

	if book == nil{
		t.Errorf("invalid book in the records")
	}

	book.title = "test*"
	book.ISBF = "test*"
	book.page = 101

	fmt.Println(book)
	if book, err := updateBookById("3d9aeb87-98fc-4008-b7fe-35f58a8a3b67", *book); book == nil && err != nil{
		t.Errorf("%v",err)
	}
}

func TestDeleteById(t *testing.T){
	if msg, err := deleteById("3d9aeb87-98fc-4008-b7fe-35f58a8a3b67"); msg == "" && err != nil{
		t.Errorf("error while deleting the book")
	}
}