package models

import (
	"testing"
	"fmt"
)

func TestAddAuthor(t *testing.T){
	fmt.Println("test author")
	tmp := Author{
		full_name: "test test",
	}
	message, err := AddAuthor(tmp)

	if message == "" && err != nil {
		t.Errorf("%v", err)
	}
}