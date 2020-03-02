package controller

import (
	"fmt"
	"github.com/aymerick/raymond"
	"github.com/beeender/dailyiker/model"
)

func (blog *Blog) pageURLHelper(page string, opts *raymond.Options) interface{} {
	ctx := opts.Ctx()

	tag := valueOfMap(ctx, "tag")
	if tag == nil {
		tag = valueOfField(ctx, "Tag")
	}
	if tag == nil  || len(tag.(*model.Tag).Name) == 0  {
		if page == "1" {
			return "/"
		}
		return fmt.Sprintf(`/page/%s/`, page)
	}

	tagName := tag.(*model.Tag).Name
	if page == "1" {
		return fmt.Sprintf("/tag/%s/", tagName)
	}
	return fmt.Sprintf(`/tag/%s/page/%s/`,tagName, page)
}
