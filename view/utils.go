package view

import "reflect"

func valueOfField(d interface{}, field string) interface{} {
	v := reflect.ValueOf(d)
	if v.Kind() != reflect.Struct {
		return nil
	}
	fv := v.FieldByName(field)
	if fv.IsValid() {
		return fv.Interface()
	}
	return nil
}

func valueOfMap(d interface{}, key string) interface{} {
	v := reflect.ValueOf(d)
	if v.Kind() != reflect.Map {
		return nil
	}
	m := v.Interface().(map[string]interface{})
	return m[key]
}
