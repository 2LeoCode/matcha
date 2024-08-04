package database

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	Id       *string `sql_name:"id" sql_type:"UUID"`
	Username *string `sql_name:"username" sql_type:"VARCHAR(16)"`
	Password *string `sql_name:"password" sql_type:"VARCHAR(16)"`
	Email    *string `sql_name:"email" sql_type:"VARCHAR(254)"`
}

type UserQueryFilter struct {
	QueryFilter
	Id       bool
	Username bool
	Password bool
	Email    bool
}

func (this *DatabaseManager) FindUsers(user *User, qf *UserQueryFilter) ([]*User, error) {
	query := "SELECT "
	userType := reflect.TypeOf(User{})
	qfType := reflect.TypeOf(UserQueryFilter{})
	if qf == nil {
		query += "*"
	} else {
		v := reflect.ValueOf(*qf)
		fieldNames := []string{}
		for i := 0; i < v.NumField(); i++ {
			field := qfType.Field(i)
			value := v.Field(i)
			if value, ok := value.Interface().(bool); ok {
				if value == true {
					name, ok := userType.FieldByName(field.Name)
					if !ok {
						return nil, errors.New("Invalid query filter (must contain all user property names mapped with booleans)")
					}
					fieldNames = append(fieldNames, name.Tag.Get("sql_name"))
				}
			} else {
				return nil, errors.New("Invalid query filter (must contain all user property names mapped with booleans)")
			}
		}
		query += strings.Join(fieldNames, ", ")
	}
	query += " FROM public.users"
	if user != nil && !reflect.DeepEqual(*user, User{}) {
		v := reflect.ValueOf(*user)
		query += " WHERE "
		conditions := []string{}
		for i := 0; i < v.NumField(); i++ {
			field := userType.Field(i)
			value := v.Field(i)

			conditions = append(conditions, fmt.Sprintf("%s = %v", field.Tag.Get("sql_name"), value.Interface()))
		}
		query += strings.Join(conditions, " AND ")
	}
	result := []*User{}
	if res, err := this.pool.Query(context.Background(), query); err != nil {
		values, err := res.Values()
		if err != nil {
			return nil, err
		}
		for value := range values {
			result := append(result, new(User))
			resultValue := reflect.ValueOf(&result[len(result)-1])
			t := reflect.TypeOf(value)
			v := reflect.ValueOf(value)
			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)
				for i := 0; i < userType.NumField(); i++ {
					userField := userType.Field(i)
					if field.Name == userField.Tag.Get("sql_name") {
						if field.Type != userField.Type.Elem() {
							return nil, fmt.Errorf("Invalid type for field %s: %s", field.Name, field.Type.Name)
						}

					}
				}
			}
		}
	}
}
