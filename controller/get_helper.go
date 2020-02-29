package controller

import (
	"github.com/aymerick/raymond"
	"strconv"
)

func (blog *Blog) getHelper(res string, opts *raymond.Options) interface{} {
	result := ""
	limit := opts.HashStr("limit")
	count := 0
	count, _ = strconv.Atoi(limit)
	ctx := opts.Ctx().(map[string]interface{})

	if res == "tags" {
		tags := blog.query.Tags(count)
		blog.completeTagsInf(tags)
		ctx["tags"] = tags
		result += opts.FnWith(ctx)
	}

	return result
}
