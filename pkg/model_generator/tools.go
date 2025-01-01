package main

import (
	"errors"
	"reflect"
)

var ErrIllegalSakeName = errors.New("非法snake name")

func SnakeToUpperCamel(snake string) (string, error) {
	lowerCamel, err := SnakeToLowerCamel(snake)
	if err != nil {
		return "", err
	}
	bytes := []byte(lowerCamel)
	bytes[0] -= 32
	return string(bytes), nil
}

func SnakeToLowerCamel(snake string) (string, error) {
	bytes := []byte(snake)
	for bytes[0] == 95 {
		bytes = bytes[1:]
	}
	length := len(bytes)
	j := 0
	for i := 0; i < length; i++ {
		if bytes[i] < 97 && bytes[i] != 95 || bytes[i] > 122 {
			return "", ErrIllegalSakeName
		}
		if bytes[i] == 95 && i+1 < length && bytes[i+1] == 95 {
			return "", ErrIllegalSakeName
		}
		if bytes[i] == 95 {
			i++
			if i == length {
				return string(bytes[:j]), nil
			}
			bytes[j] = bytes[i] - 32
		} else {
			bytes[j] = bytes[i]
		}
		j++
	}
	return string(bytes[:j]), nil
}

func QueryResultToTable(name string, queryResults []map[string]string) Table {
	table := Table{}
	table.Cols = make([]Column, 0)
	for _, row := range queryResults {
		column := Column{}
		valueOfSrc := reflect.ValueOf(&column).Elem()
		for _, field := range getFields(column) {
			val := row[field]
			f := valueOfSrc.FieldByName(field)
			v := reflect.ValueOf(val)
			if f.Type() == v.Type() {
				f.Set(v)
			}
		}
		table.Cols = append(table.Cols, column)
	}
	table.Name = name
	return table
}

func getFields(src any) []string {
	fields := make([]string, 0)

	typeOfSrc := reflect.TypeOf(src)
	for i := range typeOfSrc.NumField() {
		fields = append(fields, typeOfSrc.Field(i).Name)
	}
	return fields
}
