package models

import (
	"github.com/aymane-smi/api-test/utils"
	"fmt"
	"errors"
	//"go.uber.org/zap"
)


type Author struct{
	id int
	full_name string

}

// sampleJSON := []byte(`{
// 	"level" : "info",
// 	"encoding": "json",
// 	"outputPaths":["log.log"],
// 	"errorOutputPaths":["log.log"],
// 	"encoderConfig": {
// 		"messageKey":"message",
// 		"levelKey":"level",
// 		"levelEncoder":"lowercase"
// 	}
// }`)

// var cfg zap.Config
 
//    if err := json.Unmarshal(sampleJSON, &cfg); err != nil {
//        panic(err)
//    }
 
//    logger, err := cfg.Build()
 
//    if err != nil {
//        panic(err)
//    }
//    defer logger.Sync()

func AddAuthor(a Author) (string, error){
	db := utils.GetInstance()
	stmt, err := db.Prepare("INSERT INTO authors(full_name) VALUES($1)")
	if err != nil{
		//logger.Warn(err)
		return "", err
	}
	if _, err := stmt.Exec(a.full_name); err != nil{
		//logger.Warn(err)
		return "", err
	}
	return "new row inserted in authors", nil

}

func UpdateAuthor(a Author) (*Author, error){
	db := utils.GetInstance()
	stmt, err := db.Prepare("UPDATE authors SET full_name= $1 WHERE id = $2")
	if err != nil {
		//logger.Warn(err)
		return nil, err
	}

	if _, err := stmt.Exec(a.full_name, a.id); err != nil{
		return nil, err
	}

	return &a, err
}

func GetAuthorById(id int) *Author{
	db := utils.GetInstance()
	stmt, err := db.Prepare("SELECT * FROM authors WHERE id = $1")
	if err != nil{
		//logger.Warn(err)
		return nil
	}

	row := stmt.QueryRow(id)
	var tmp_author Author
	if err := row.Scan(&tmp_author.id, &tmp_author.full_name); err != nil{
		//logger.Warn(err)
		return nil
	}
	return &tmp_author
	
}

func DeleteAuthorById(id int) (string, error){
	db := utils.GetInstance()
	stmt, err := db.Prepare("DELETE FROM authors WHERE id = $1")
	if err != nil{
		fmt.Println("prepare error")
		//logger.Warn(err)
		return "", err
	}
	x, err := stmt.Exec(id);
	rowsAffected, err := x.RowsAffected()
	if rowsAffected == 0 || err != nil{
		//logger.Warn(err)
		return "", errors.New("invalid id")
	}

	return fmt.Sprintf("book with id '%d' was deleted", id), nil
}