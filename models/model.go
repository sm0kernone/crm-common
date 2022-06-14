package models

import (
	"context"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	Type       = "type"
	Label      = "label"
	SQL        = "sql"
	Group      = "group"
	JSON       = "json"
	Related    = "related"
	Name       = "name"
	Values     = "values"
	Order      = "order"
	Edit       = "edit"
	Questioner = "questioner"
	View       = "view"
)

var keys = map[string]func(field reflect.StructField) interface{}{
	Name: func(field reflect.StructField) interface{} {

		var json = strings.Split(GetValueByTag(field, JSON), ",")

		if len(json[0]) <= 0 {
			return field.Type.Name()
		}

		return json[0]
	},
	Type: func(field reflect.StructField) interface{} {

		var types = strings.Split(GetValueByTag(field, Type), ",")

		if len(types[0]) <= 0 {
			types[0] = field.Type.Name()
		}

		return types
	},
	Label: func(field reflect.StructField) interface{} {

		var label = GetValueByTag(field, Label)

		if len(label) <= 0 {
			return field.Name
		}

		return label
	},
	Group: func(field reflect.StructField) interface{} {

		var groups = strings.Split(GetValueByTag(field, Group), ",")

		if len(groups[0]) <= 0 {
			return []string{}
		}

		return groups
	},
	Values: func(field reflect.StructField) interface{} {

		var groups = strings.Split(GetValueByTag(field, Values), ",")

		if len(groups[0]) <= 0 {
			return []string{}
		}

		return groups
	},
	Edit: func(field reflect.StructField) interface{} {

		var groups = strings.Split(GetValueByTag(field, Edit), ",")

		if len(groups[0]) <= 0 {
			return []string{"ALL"}
		}

		return groups
	},
	Related: func(field reflect.StructField) interface{} {

		var label = GetValueByTag(field, Related)

		if len(label) <= 0 {
			return nil
		}

		return label
	},
	Order: func(field reflect.StructField) interface{} {
		var o = GetValueByTag(field, Order)

		if len(o) <= 0 {
			return 0
		}

		order, e := strconv.Atoi(o)

		if e != nil {
			return 0
		}

		return order
	},
	Questioner: func(field reflect.StructField) interface{} {
		var o = GetValueByTag(field, Questioner)

		if len(o) <= 0 {
			return 0
		}

		order, e := strconv.Atoi(o)

		if e != nil {
			return 0
		}

		return order
	},
	View: func(field reflect.StructField) interface{} {
		var view = GetValueByTag(field, View)

		if len(view) <= 0 {
			return field.Name
		}

		return view
	},
}

func GetModelInfo(model interface{}) map[string]map[string]interface{} {

	var ref = reflect.TypeOf(model)

	var result = map[string]map[string]interface{}{}

	for i := 0; i < ref.NumField(); i++ {

		var field = ref.Field(i)

		var key = strings.Split(GetValueByTag(field, "json"), ",")[0]

		if len(key) <= 0 || key == "-" {
			continue
		}

		result[key] = map[string]interface{}{}

		for tag, call := range keys {
			result[key][tag] = call(field)
		}
	}

	return result
}

func GetValueByTag(f reflect.StructField, tag string) string {
	return f.Tag.Get(tag)
}

// Base contains common table fields
type Base struct {
	ID        int `sql:"id" json:"id"`
	CreatedAt int `sql:"created_at" json:"created_at,omitempty" label:"Created At" type:"datetime" structs:"-"`
	UpdatedAt int `sql:"updated_at" json:"updated_at,omitempty" label:"Updated At" type:"datetime" structs:"-"`
	DeletedAt int `sql:"deleted_at" json:"deleted_at,omitempty" label:"Deleted At" type:"datetime" structs:"-"`
}

// BeforeUpdate hooks into update operations
func (p *Base) BeforeUpdate(ctx context.Context) (context.Context, error) {
	p.UpdatedAt = int(time.Now().UTC().Unix())

	return ctx, nil
}
