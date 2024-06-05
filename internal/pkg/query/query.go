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

	fmt.Printf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), strings.Join(values, ", "))
	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, strings.Join(columns, ", "), strings.Join(values, ", "))
}

func ForUpdate(entity TableInterface) string {
	entityValue := reflect.ValueOf(entity)
	entityType := reflect.TypeOf(entity)
	tableName := entity.GetTableName()

	var columns, values []string
	for i := 0; i < entityType.NumField(); i++ {
		field := entityType.Field(i)
		fieldValue := entityValue.Field(i)
		if isZero(fieldValue) {
			continue
		}

		columns = append(columns, field.Tag.Get("db"))
		values = append(values, fmt.Sprintf(":%v", fieldValue.Interface()))
	}

	return fmt.Sprintf("UPDATE %s SET %s WHERE id = :id", tableName, strings.Join(columns, ", "), reflect.ValueOf(entity).Field(0))
}

func isZero(value reflect.Value) bool {
	zero := reflect.Zero(value.Type()).Interface()
	return reflect.DeepEqual(value.Interface(), zero)
}

type TableInterface interface {
	GetTableName() string
}
