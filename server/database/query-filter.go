package database

import (
	"errors"
	"matcha/utils"
	"reflect"
)

type QueryFilter[T any] struct {
	target T
}

func NewQueryFilter[T any](target T) *QueryFilter[T] {
	return &QueryFilter[T]{target}
}

func (this *QueryFilter) String() (string, error) {
	objType := reflect.TypeOf(this)

	if objType.Kind() != reflect.Struct {
		return "", errors.New("obj must be a struct")
	}

	if queryFilter == nil {
		return "*", nil
	}

	queryFilterValue := reflect.ValueOf(queryFilter).Elem()
	queryFilterType := queryFilterValue.Type()

	if queryFilterType.Kind() != reflect.Struct {
		return "", errors.New("queryFilter must be a pointer to a struct")
	}

	if sameFieldNames, _ := utils.HaveSameFieldNames(obj, *queryFilter); !sameFieldNames {
		return "", errors.New("obj and the value pointed by queryFiler must have the same field names")
	}

	fields := []string{}
	for i := 0; i < .NumField(); i++ {

	}
}
