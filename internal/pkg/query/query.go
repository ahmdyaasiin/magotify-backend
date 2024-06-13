package query

import (
	"fmt"
	"reflect"
	"strings"
)

func ForCreate(entity TableInterface) string {
	entityValue := reflect.ValueOf(entity).Elem()
	entityType := entityValue.Type()
	tableName := entity.GetTableName()

	var columns, values []string
	for i := 0; i < entityType.NumField(); i++ {
		field := entityType.Field(i)
		fieldValue := entityValue.Field(i)
		if isZero(fieldValue) {
			continue
		}

		columns = append(columns, field.Tag.Get("db"))
		values = append(values, fmt.Sprintf(":%v", field.Tag.Get("db")))
	}

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), strings.Join(values, ", "))
}

func ForUpdate(entity TableInterface) string {
	entityValue := reflect.ValueOf(entity).Elem() // Use Elem() to get the underlying struct
	entityType := entityValue.Type()
	tableName := entity.GetTableName()

	var columns []string
	for i := 0; i < entityType.NumField(); i++ {
		field := entityType.Field(i)
		fieldValue := entityValue.Field(i)
		if isZero(fieldValue) {
			continue
		}

		columnName := field.Tag.Get("db")
		columns = append(columns, fmt.Sprintf("%s = :%s", columnName, columnName))
	}

	return fmt.Sprintf("UPDATE %s SET %s WHERE id = :id", tableName, strings.Join(columns, ", "))
}

func ForDelete(entity TableInterface) string {
	tableName := entity.GetTableName()
	return fmt.Sprintf("DELETE FROM %s WHERE id = :id", tableName)
}

func isZero(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String, reflect.Array:
		return value.Len() == 0
	case reflect.Map, reflect.Slice, reflect.Chan, reflect.Ptr, reflect.Interface:
		return value.IsNil()
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Complex64, reflect.Complex128:
		return value.Complex() == 0
	case reflect.Struct:
		// untuk struct, periksa apakah semua field bernilai nol
		for i := 0; i < value.NumField(); i++ {
			if !isZero(value.Field(i)) {
				return false
			}
		}
		return true
	}
	return false
}

type TableInterface interface {
	GetTableName() string
}
