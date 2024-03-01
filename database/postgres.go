package database

import (
	"os"
	"github.com/go-pg/pg/v10"
)

func DBConn() (*pg.DB, error)  {
	addr := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	userName := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	
	db := pg.Connect(&pg.Options{
		Addr:     addr + ":" + port,
		User:     userName,
		Password: password,
		Database: database,
	})

	if err := checkConn(db); err != nil {
		return nil, err
	}
	return  db, nil;
}

func checkConn(db *pg.DB) error {
	var n int
	_, err := db.QueryOne(pg.Scan(&n), "SELECT 1")
	return  err
}