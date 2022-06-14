package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func InitSqliteClient(dbName string) (*sql.DB, error) {
	os.Remove(dbName)

	return sql.Open("sqlite3", dbName)
}
