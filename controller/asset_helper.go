package controller

import (
	"github.com/aymerick/raymond"
)

func (blog *Blog) assetHelper(context interface{}, _ *raymond.Options) interface{} {
	return "/assets/" + raymond.Str(context)
}

