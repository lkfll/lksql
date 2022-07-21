package entity

import "github.com/lkfll/lksql"

var FieldConfigRepository = lksql.EntityDefaultFacory(FieldConfig{})

type FieldConfig struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
}
