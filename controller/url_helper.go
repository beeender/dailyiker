package controller

import (
	"github.com/aymerick/raymond"
	"net/url"
)

func (blog *Blog) urlHelper(opts *raymond.Options) interface{} {
	absolute := false
	if opts.Hash() != nil {
		absolute = raymond.IsTrue(opts.HashProp("absolute"))
	}

	ctx := opts.Ctx()
	if ctx == nil {
		return "/"
	}
	v := valueOfField(ctx, "URL")
	if v == nil {
		v = valueOfMap(ctx, "url")
	}
	if v == nil {
		return "/"
	}
	urlStr := raymond.Str(v)

	if !absolute {
		return raymond.SafeString(joinPath(blog.Config.URLPrefix, urlStr))
	}

	u, e := url.Parse(urlStr)
	if e != nil {
		return "/"
	}
	if u.IsAbs() {
		return raymond.SafeString(urlStr)
	}
	tmpUrl := blog.RootURL
	tmpUrl.Path = joinPath(tmpUrl.Path, u.Path)
	return raymond.SafeString(tmpUrl.String())
}
