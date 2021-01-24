package config

import (
	"flag"
	"log"
	"os"
)

type Config struct {
	HTTPAddr    string
	UserLogFile string
	SQLiteFile  string
}

func MakeConfig() Config {
	var c Config

	HTTPAddr, found := os.LookupEnv("HTTP_ADDR")
	if !found {
		HTTPAddr = ":8080"
	}
	flag.StringVar(&c.HTTPAddr, "http.addr", HTTPAddr, "HTTP listen address")

	UserLogFile, found := os.LookupEnv("USER_LOG_FILE")
	if !found {
		path, err := os.Getwd()
		if err != nil {
			log.Printf("Not found current dir: %s", err)
		}
		UserLogFile = path + "/user.log"
	}
	flag.StringVar(&c.UserLogFile, "user.log.file", UserLogFile,
		"Path and file name for user logging")

	SQLiteFile, found := os.LookupEnv("SQLITE_FILE")
	if !found {
		SQLiteFile = ""
	}
	flag.StringVar(&c.SQLiteFile, "sqlite.file", SQLiteFile,
		"Path and file name for sqlite")

	flag.Parse()

	return c
}
