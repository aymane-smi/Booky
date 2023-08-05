package utils

import (
	"encoding/json"

	"go.uber.org/zap"
)

var Log *zap.Logger

func InitLogger(){
	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["./logs/api.log"],
		"errorOutputPaths": ["./logs/api.log"],
		"initialFields": {"foo": "bar"},
		"encoderConfig": {
		  "timeKey": "logged at", 
		  "timeEncoder": "ISO8601",
		  "messageKey": "message",
		  "levelKey": "level",
		  "levelEncoder": "lowercase"
		}
	  }`)
  
	  var cfg zap.Config
	  if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		  panic(err)
	  }
	  logger := zap.Must(cfg.Build())
	  defer logger.Sync()
  
	  logger.Info("logger construction succeeded")

	  Log = logger
}