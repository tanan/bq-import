package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/tanan/bq-import/config"
)

type SQLHandler struct {
	DB *sql.DB
}

func NewSqlHandler(config config.DBConfig) *SQLHandler {
	db, err := sql.Open("postgres", config.GetConnInfo())
	if err != nil {
		fmt.Errorf("database connection error. %v", err)
	}
	return &SQLHandler{
		DB: db,
	}
}
