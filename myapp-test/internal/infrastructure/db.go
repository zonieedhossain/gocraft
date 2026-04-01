package infrastructure

import (
	"fmt"
	"os"

	
	"database/sql"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	
)


func NewDB() (*sql.DB, error) {
	var sqldb *sql.DB
	var err error

	switch os.Getenv("DB_TYPE") {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
		sqldb, err = sql.Open("mysql", dsn)
	case "sqlite":
		sqldb, err = sql.Open("sqlite3", os.Getenv("DB_NAME")+".db")
	default: // postgres
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
			os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"))
		sqldb, err = sql.Open("postgres", dsn)
	}

	return sqldb, err
}

