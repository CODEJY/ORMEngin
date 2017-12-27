package dao

import (
	"errors"
	"reflect"
	"strings"
)

/*
helpful func about Insert
*/
// generate insert statement

func insertStmt(o interface{}) (string, error) {
	tableName, err := getTableName(o)
	if err != nil {
		return "", err
	}
	stmt := "INSERT " + tableName + " SET "
	fields, _, err := getTableField(o)
	if err != nil {
		return "", err
	}

	for i := 0; i < len(fields)-1; i++ {
		stmt += fields[i] + "=?,"
	}
	stmt += fields[len(fields)-1] + "=?"

	return stmt, nil
}

// get database table name
func getTableName(o interface{}) (string, error) {
	t := reflect.TypeOf(o)
	if t.Name() == "" {
		return "", errors.New("non-exist interface type")
	}
	return strings.ToLower(t.Name()), nil
}

// get table field's name and value
func getTableField(o interface{}) ([]string, []interface{}, error) {
	fieldNames := make([]string, 0)
	fieldValues := make([]interface{}, 0)

	s := reflect.ValueOf(o)
	typeOfO := s.Type()
	if typeOfO.Kind() != reflect.Struct {
		return []string{}, []interface{}{}, errors.New("no struct type")
	}
	for i := 0; i < s.NumField(); i++ {
		fieldNames = append(fieldNames, strings.ToLower(typeOfO.Field(i).Name))
		fieldValues = append(fieldValues, s.Field(i).Interface())
	}

	return fieldNames, fieldValues, nil
}

// generate query statement
func queryStmt(tableName string) (string, error) {
	if tableName == "" {
		return "", errors.New("non-exist interface type")
	}

	return "SELECT * FROM " + strings.ToLower(tableName), nil
}
