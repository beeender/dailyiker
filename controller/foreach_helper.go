package controller

import (
	"github.com/aymerick/raymond"
	"reflect"
)

// See https://ghost.org/docs/api/v3/handlebars-themes/helpers/foreach/
func (blog *Blog) foreachHelper(context interface{}, options *raymond.Options) interface{} {
	if !raymond.IsTrue(context) {
		return options.Inverse()
	}

	result := ""
	val := reflect.ValueOf(context)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			/* This would trigger a null pointer error. No idea why yet.
			dataFrame := raymond.NewDataFrame()
			dataFrame.Set("index", i)
			dataFrame.Set("number", i + 1)
			dataFrame.Set("first", i == 0)
			dataFrame.Set("last", i == val.Len() - 1)
			dataFrame.Set("odd", i % 2 == 1)
			dataFrame.Set("even", i % 2 == 0)
			// TODO: what to do with rowStart and rowEnd?
			result += options.FnCtxData(val.Index(i).Interface(), dataFrame)
			*/
			result += options.FnWith(val.Index(i).Interface())
		}
	}
	return raymond.SafeString(result)
}

