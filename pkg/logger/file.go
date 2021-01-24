package logger

import (
	"log"
	"os"
)

var (
	Log *log.Logger
)

func File(filePath string) {

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Print(err)
	}
	Log = log.New(file, "", 0)
}
