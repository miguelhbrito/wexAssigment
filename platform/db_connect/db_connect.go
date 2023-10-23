package dbconnect

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InitDB(DBDriver, DBSource string) (db *sql.DB) {
	db, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		fmt.Println("error to open db", err)
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("error to ping db", err)
		panic(err)
	}
	fmt.Println("Successfully connected!")
	return db
}
