package database

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func createTable(pool *pgxpool.Pool, schema, table string, obj any) error {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.%s (", schema, table)
	t := reflect.TypeOf(obj)
	fields := []string{}
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Kind() != reflect.Ptr {
			return errors.New("A table type must contain only pointer fields")
		}
		fields = append(fields, fmt.Sprintf("%s %s", field.Tag.Get("sql_name"), field.Tag.Get("sql_type")))
	}
	query += fmt.Sprintf("%s)", strings.Join(fields, ", "))
	_, err := pool.Exec(context.Background(), query)
	return err
}

func createPublicTable(pool *pgxpool.Pool, table string, obj any) error {
	return createTable(pool, "public", table, obj)
}
