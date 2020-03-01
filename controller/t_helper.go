package controller

import (
	"github.com/aymerick/raymond"
)

func (blog *Blog) tHelper(item string, opts *raymond.Options) interface{} {
	return item
}
