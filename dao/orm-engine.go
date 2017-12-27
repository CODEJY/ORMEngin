package dao

import (
	"database/sql"

	"github.com/CODEJY/ORMEngine/sqlt"

	_ "github.com/go-sql-driver/mysql"
)

// ORMEngine definition
type ORMEngine struct {
	sqlt.SQLTemplate
}

// NewEngine create a new engine for database operation
func NewEngine(driverName, dataSourceName string) *ORMEngine {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	engine := &ORMEngine{sqlt.NewSQLTemplate(db)}

	return engine
}
