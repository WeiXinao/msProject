package main

import (
	"fmt"
	"strings"
)

type Column struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
	Comment string
}

func (c Column) ToModelType() string {
	switch {
	case strings.Contains(c.Type, "bigint"):
		return "int64"
	case strings.Contains(c.Type, "varchar"):
		return "string"
	case strings.Contains(c.Type, "text"):
		return "string"
	case strings.Contains(c.Type, "tinyint"):
		return "int"
	case strings.Contains(c.Type, "int") &&
		!strings.Contains(c.Type, "tinyint") &&
		!strings.Contains(c.Type, "bigint"):
		return "int"
	case strings.Contains(c.Type, "double"):
		return "float64"
	default:
		return ""
	}
}

func (c Column) ToMessageType() string {
	switch {
	case strings.Contains(c.Type, "bigint"):
		return "int64"
	case strings.Contains(c.Type, "varchar"):
		return "string"
	case strings.Contains(c.Type, "text"):
		return "string"
	case strings.Contains(c.Type, "tinyint"):
		return "int32"
	case strings.Contains(c.Type, "int") &&
		!strings.Contains(c.Type, "tinyint") &&
		!strings.Contains(c.Type, "bigint"):
		return "int32"
	case strings.Contains(c.Type, "double"):
		return "double"
	default:
		return ""
	}
}

type Table struct {
	Name string
	Cols []Column
}

func (t Table) ToModel() (Model, error) {
	t.Name = strings.TrimSpace(t.Name)
	model := Model{}
	name, err := SnakeToUpperCamel(t.Name)
	if err != nil {
		return Model{}, err
	}
	model.Name = name
	model.TableName = t.Name

	model.Fields = make([]Field, 0)
	for _, col := range t.Cols {
		field := Field{}
		field.Field, err = SnakeToUpperCamel(col.Field)
		if err != nil {
			return Model{}, err
		}
		field.Column = fmt.Sprintf("'%s'", col.Field)
		field.Type = col.ToModelType()
		field.DBType = col.Type

		if col.Null == "YES" {
			field.Null = "null"
		} else {
			field.Null = "notnull"
		}

		if col.Key != "" {
			if col.Key == "PRI" {
				field.Key = "pk"
			}
		}

		if field.Type == "string" {
			field.Default = fmt.Sprintf("default('%s')", col.Default)
		} else {
			if col.Default != "" {
				field.Default = fmt.Sprintf("default(%s)", col.Default)
			}
		}

		if col.Extra == "auto_increment" {
			field.Extra = "autoincr"
		} else {
			field.Extra = col.Extra
		}

		if col.Comment != "" {
			field.Comment = fmt.Sprintf("comment('%s')", col.Comment)
		}

		model.Fields = append(model.Fields, field)
	}
	return model, nil
}

type Field struct {
	Field   string
	Column  string
	Type    string
	DBType  string
	Null    string
	Key     string
	Default string
	Extra   string
	Comment string
}

type Model struct {
	Name      string
	TableName string
	Fields    []Field
}

type Message struct {
	Name   string
	Fields []Column
}

func (t Table) ToMessage() (Message, error) {
	t.Name = strings.TrimSpace(t.Name)
	message := Message{}
	name, err := SnakeToUpperCamel(t.Name)
	if err != nil {
		return Message{}, err
	}

	message.Name = name + "Message"
	message.Fields = make([]Column, 0)
	for _, col := range t.Cols {
		field := Column{}
		field.Field, err = SnakeToLowerCamel(col.Field)
		if err != nil {
			return Message{}, err
		}
		field.Type = col.ToMessageType()

		message.Fields = append(message.Fields, field)
	}
	return message, nil
}
