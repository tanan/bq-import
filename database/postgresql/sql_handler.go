package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/tanan/bq-import/config"
)

type SQLHandler struct {
	DB *sql.DB
}

func NewSqlHandler(config config.DBConfig) (*SQLHandler, error) {
	db, err := sql.Open("postgres", config.GetConnInfo())
	if err != nil {
		return nil, fmt.Errorf("database connection error. %v", err)
	}
	return &SQLHandler{
		DB: db,
	}, nil
}

func (s *SQLHandler) GetRow() {
	// TODO implementation
	fmt.Print("TODO implement")
}
