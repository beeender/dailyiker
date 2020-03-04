package controller

import (
	"github.com/aymerick/raymond"
)

func (blog *Blog) assetHelper(context interface{}, _ *raymond.Options) interface{} {
	return joinPath(blog.Config.URLPrefix, "/assets/" + raymond.Str(context))
}
