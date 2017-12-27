package dao

import (
	"database/sql"
	"reflect"

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

// query all the entries of the table
func (engine *ORMEngine) Find(o interface{}) error {
	sliceValue := reflect.Indirect(reflect.ValueOf(o))

	sliceElementType := sliceValue.Type().Elem()
	data := sliceElementType.Elem()
	queryString, err := queryStmt(data.Name())
	if err != nil {
		return err
	}

	rows, err := engine.Query(queryString)
	if err != nil {
		return err
	}

	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	newSlice := reflect.MakeSlice(sliceValue.Type(), 0, 4)

	for rows.Next() {
		pv := reflect.New(data)
		fieldAdr := pv.Elem()

		for i := 0; i < fieldAdr.NumField(); i++ {
			f := fieldAdr.Field(i)
			values[i] = f.Addr().Interface()
		}

		rows.Scan(values...)

		newSlice = reflect.Append(newSlice, pv)
	}

	s := reflect.ValueOf(o).Elem()
	s.Set(newSlice)

	return nil
}
