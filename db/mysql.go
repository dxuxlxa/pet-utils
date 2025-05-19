package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

// Example usage:
//
// db, err := MysqlConnector(mySqlArgs)
//
// defer db.Close()
//
// ... use db ...
//
// _, err = db.Exec(createTableQuery) //  Table creation moved out
func MysqlConnector(args SqlFields) (*sql.DB, error) {
	var err error
	var db *sql.DB
	// mysql connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", args.Username, args.Password, args.Host, args.Port, args.Dbname)
	for range 10 {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				log.Printf("MySql Connection %s for %s service completed!\n", args.Dbname, args.Service)
				return db, nil
			}
		}
		log.Printf("Database connection failed: %v. Retrying in 5 seconds...\n", err)
		time.Sleep(5 * time.Second)
	}
	return nil, fmt.Errorf("could not connect to the database: %w", err)

}
