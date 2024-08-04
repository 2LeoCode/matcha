package utils

import (
	"errors"
	"reflect"
)

func HaveSameFieldNames(lhs, rhs any) (bool, error) {
	lhsType := reflect.TypeOf(lhs)
	rhsType := reflect.TypeOf(rhs)

	if lhsType.Kind() != reflect.Struct || rhsType.Kind() != reflect.Struct {
		return false, errors.New("Both parameters aren't structs")
	}
	for i := 0; i < lhsType.NumField(); i++ {
		if _, ok := rhsType.FieldByName(lhsType.Field(i).Name); !ok {
			return false, nil
		}
	}
	for i := 0; i < rhsType.NumField(); i++ {
		if _, ok := lhsType.FieldByName(rhsType.Field(i).Name); !ok {
			return false, nil
		}
	}
	return true, nil
}
