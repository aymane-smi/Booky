package models

import (
	"fmt"
	"testing"

	"github.com/aymane-smi/api-test/utils"
)

func TestGetById(t *testing.T){
	utils.InitLogger()
	book := GetBookById("test_id")
	if book == nil {
		t.Errorf("error bookis empty");
	}
}

func TestAddBook(t *testing.T){
	utils.InitLogger()
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

func TestAddBookTesting(t *testing.T){
	utils.InitLogger()
	tmp := Book{
		ISBF: "test",
		Title: "test title",
		Page: 24,
		Author: 1,
	}
	message, err := AddBookTest(tmp)

	if message == "" && err != nil {
		t.Errorf("%v", err)
	}
}

func TestUpdateBook(t *testing.T){
	utils.InitLogger()
	book := GetBookById("test_id")

	if book == nil{
		t.Errorf("invalid book in the records")
	}else{
		book.Title = "test*"
		book.ISBF = "test*"
		book.Page = 101

		fmt.Println(book)
		if book, err := UpdateBook(*book); book == nil && err != nil{
			t.Errorf("%v",err)
		}
	}
}

func TestDeleteById(t *testing.T){
	if msg, err := DeleteById("test_id"); msg == "" && err != nil{
		t.Errorf("error while deleting the book")
	}
}