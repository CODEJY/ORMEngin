package dao

import (
	"database/sql"

	"github.com/CODEJY/ORMEngine/sqlt"

	_ "github.com/go-sql-driver/mysql"
)

// ORMEngine struct definition
type ORMEngine struct {
	sqlt.SQLTemplate
}

// create a new engine
func NewEngine(driverName, dataSourceName string) *ORMEngine {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	engine := &ORMEngine{sqlt.NewSQLTemplate(db)}
	return engine
}

// Insert new data entry into the table
func (e *ORMEngine) Insert(o interface{}) (int, error) {
	insertQuery, err := insertStmt(o)
	if err != nil {
		return 0, err
	}
	_, args, err := getTableField(o)
	if err != nil {
		return 0, err
	}

	_, err = e.Exec(insertQuery, args...)
	if err != nil {
		return 0, err
	}

	return 1, nil
}
