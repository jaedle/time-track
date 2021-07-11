package persistence

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase(dataSource string) (*sql.DB, error) {
	return sql.Open("mysql", dataSource)
}


