package controller

import (
	"github.com/aymerick/raymond"
	"strconv"
)

func (blog *Blog) getHelper(conditional interface{}, opts *raymond.Options) interface{} {
	if !raymond.IsTrue(conditional) {
		return opts.Inverse()
	}
	res := conditional.(string)

	result := ""
	limit := opts.HashStr("limit")
	count := 0
	count, _ = strconv.Atoi(limit)

	ctxMap := make(map[string]interface{})

	if res == "tags" {
		tags := blog.query.Tags(count)
		blog.completeTagsInf(tags)
		ctxMap["tags"] = tags
		result += opts.FnWith(ctxMap)
	}

	return result
}
