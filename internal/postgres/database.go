package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDatabase(dbConnString string) *sql.DB {
	db, err := sql.Open("postgres", dbConnString)
	if err != nil {
		panic(fmt.Sprintf("Can't connect to the database. %v", err))
	} else {
		fmt.Println("Database Connected.")
	}
	db.SetMaxOpenConns(75) // The RDS instance has a max of 75 open connections
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(300)
	return db
}
