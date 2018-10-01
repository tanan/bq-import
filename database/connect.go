package database

import (
	"github.com/tanan/bq-import/config"
	"github.com/tanan/bq-import/database/postgresql"
)

func Connect(config config.DBConfig) (SQLHandler, error) {
	switch config.SQLType {
	case "postgres":
		return postgresql.NewSqlHandler(config)
	default:
		return postgresql.NewSqlHandler(config)
	}
	return nil, nil
}
