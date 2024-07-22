package utils

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

func InitDB(user, password, dbname, host, port string) (*sql.DB, error) {
	p, err := strconv.Atoi(port)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, p, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("сonnected to the database")
	return db, nil
}
