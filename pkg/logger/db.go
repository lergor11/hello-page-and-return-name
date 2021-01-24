package logger

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func SQLite(SQLitePath string, name string, data string) {
	db, err := sql.Open("sqlite3", SQLitePath)
	if err != nil {
		log.Printf("Error open db file: %s", err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS user (name TEXT, data DATA)")
	if err != nil {
		log.Printf("Error create table in db: %s", err)
	}

	_, err = db.Exec("insert into user (name, data) values ($1, $2)", name, data)
	if err != nil {
		log.Printf("Error insert into db: %s", err)
	}

}
