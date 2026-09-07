package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:31415926@tcp(127.0.0.1:3306)/healthtrack")
	if err != nil {
		return nil, fmt.Errorf("Error opening database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Failed to connect database: %w", err)
	}

	fmt.Println("Successfully connected to database :D")
	return db, nil
}
