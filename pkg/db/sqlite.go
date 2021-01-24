package db

import (
	"database/sql"
	"log"
)

func GetUser(SQLitePath string, name string) (userData []string) {
	db, err := sql.Open("sqlite3", SQLitePath)
	if err != nil {
		log.Printf("Error open db file: %s", err)
		return []string{}
	}
	defer db.Close()

	selectUser, err := db.Query("select data from user where name = $1", name)
	if err != nil {
		log.Printf("Error select into db: %s", err)
		return []string{}
	}
	defer selectUser.Close()
	var p string
	for selectUser.Next() {
		err := selectUser.Scan(&p)
		if err != nil {
			log.Printf("Error parse object: %s", err)
			continue
		}
		userData = append(userData, p)
	}
	return userData
}
