package controller

import (
	"github.com/aymerick/raymond"
)

func (blog *Blog) imgURLHelper(url string, _ *raymond.Options) interface{} {
	//size := options.HashStr("size")
	return url
}
