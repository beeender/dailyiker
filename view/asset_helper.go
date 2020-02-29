package view

import "github.com/aymerick/raymond"

func (theme *Theme) assetHelper(context interface{}, _ *raymond.Options) interface{} {
	return "/assets/" + raymond.Str(context)
}

