package controller

import (
	"github.com/aymerick/raymond"
	"net/url"
)

func (blog *Blog) imgURLHelper(u string, _ *raymond.Options) interface{} {
	tmpUrl, err := url.Parse(u)
	if err != nil {
		return ""
	}
	if tmpUrl.IsAbs() {
		return u
	}

	return joinPath(blog.URLPrefix, u)
}
