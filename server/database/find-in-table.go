package database

import (
	"reflect"

	"github.com/jackc/pgx/v5/pgxpool"
)

func findInTable[T *any, TQueryFilter *queryFilter](pool *pgxpool.Pool, schema, table string, obj T, queryFilter TQueryFilter) {
	objType := reflect.TypeOf(obj).Elem()
	objValue := reflect.ValueOf(obj)
	queryFilterType := reflect.TypeOf(queryFilter).Elem()
	queryFilterValue := reflect.ValueOf(queryFilter)

	query := "SELECT "
}
