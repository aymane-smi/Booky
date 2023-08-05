package models

import (
	"fmt"
	"testing"
)

func TestAddAuthor(t *testing.T){
	tmp := Author{
		Full_name: "test test",
	}
	message, err := AddAuthor(tmp)

	if message == "" && err != nil {
		t.Errorf("%v", err)
	}
}

func TestUpdateAuthor(t *testing.T){
	tmp := Author{
		Id: 1,
		Full_name: "test* test*",
	}

	if author, err := UpdateAuthor(tmp); author == nil && err != nil{
		t.Errorf("%v", err)
	}
}

func TestGetAuthorById(t *testing.T){
	if author := GetAuthorById(1); author == nil{
		fmt.Println(author)
		t.Errorf("invalid id")
	}
}

func TestDeleteAuthorById(t *testing.T){
	if msg, err := DeleteAuthorById(3); msg == "" && err != nil{
		t.Errorf("%v", err)
	}
}