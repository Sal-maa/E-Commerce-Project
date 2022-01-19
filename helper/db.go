package helper

import (
	"database/sql"
)

func InitDB(connectionString string) (*sql.DB, error) {
	return sql.Open("mysql", connectionString)
}
