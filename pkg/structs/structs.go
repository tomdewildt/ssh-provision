package structs

import (
	"errors"
	"reflect"
)

// Map is used to convert a struct to an map of string and interface.
// It takes an struct as parameter. It returns an map of strings and
// interfaces and nil or nil and an error if one occurred.
func Map(input interface{}) (map[string]interface{}, error) {
	if input == nil || reflect.ValueOf(input).Kind() != reflect.Struct {
		return nil, errors.New("invalid input")
	}

	output := make(map[string]interface{})
	val := reflect.ValueOf(input)
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		output[typ.Field(i).Name] = val.Field(i).Interface()
	}

	return output, nil
}
