package controller

import (
	"github.com/aymerick/raymond"
	"github.com/beeender/dailyiker/model"
	"regexp"
	"strconv"
)

var (
	reTagsCount = regexp.MustCompile(`count:([><]?)(\d+)`)
)

func (blog *Blog) hasHelper(opts *raymond.Options) interface{} {
	tagsOpt := opts.HashStr("tag")
	if tagsOpt == "" {
		return opts.Inverse()
	}

	ctx := opts.Ctx()
	if !raymond.IsTrue(ctx) {
		return opts.Inverse()
	}

	v := valueOfField(ctx, "Tags")
	if v == nil {
		return opts.Inverse()
	}
	tags := v.([]model.Tag)

	matches := reTagsCount.FindStringSubmatch(tagsOpt)
	if len(matches) != 3 {
		return opts.Inverse()
	}

	fit := false
	count,_ := 	strconv.Atoi(matches[2])
	switch matches[1] {
	case "":
		if count == len(tags) {
			fit = true
		}
	case ">":
		if len(tags) > count  {
			fit = true
		}
	case "<":
		if len(tags) < count {
			fit = true
		}
	}

	if fit {
		return raymond.SafeString(opts.Fn())
	}
	return opts.Inverse()
}
