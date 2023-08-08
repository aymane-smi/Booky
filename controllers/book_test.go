package controllers

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/aymane-smi/api-test/utils"
)

func TestGetBookById(t *testing.T){

	utils.InitLogger()

	req := httptest.NewRequest("GET", "http://localhost:8000/book/21a1e439-baf5-400a-b515-c10ed7ead9b5", nil);
	w := httptest.NewRecorder()

	GetBookById(w, req)

	resp := w.Result()
	fmt.Println(resp.StatusCode)
	if resp.StatusCode != 200{
		t.Errorf("invalid request!")
	}
	// apitest.New().
	// 		HandlerFunc(GetBookById).
	// 		Get("/book/1").
	// 		Expect(t).
	// 		Status(http.StatusOK).
	// 		End()
}