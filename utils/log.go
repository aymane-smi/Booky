package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var Log *zap.Logger

func InitLogger(){
	if err := godotenv.Load(".env"); err != nil {
        fmt.Println("Error loading .env file")
    }

	var path string

	if os.Getenv("WORK_MODE") == "TEST"{
		path = "../logs/api.log"
	}else if os.Getenv("WORK_MODE") == "DEV"{
		path = "./logs/api.log"
	}
	rawJSON := []byte(fmt.Sprintf(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["%s"],
		"errorOutputPaths": ["%s"],
		"initialFields": {"foo": "bar"},
		"encoderConfig": {
		  "timeKey": "logged at", 
		  "timeEncoder": "ISO8601",
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}
	  }`, path, path))
  
	  var cfg zap.Config
	  if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		  panic(err)
	  }
	  logger := zap.Must(cfg.Build())
	  defer func(){
		err := logger.Sync()
		if err != nil{
			panic(err)
		}
	  }()
  
	  logger.Info("logger construction succeeded")

	  Log = logger
}