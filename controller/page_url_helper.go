package controller

import (
	"fmt"
	"github.com/aymerick/raymond"
	"github.com/beeender/dailyiker/model"
)

func (blog *Blog) pageURLHelper(page string, opts *raymond.Options) interface{} {
	ctx := opts.Ctx()
	url := ""

	tag := valueOfMap(ctx, "tag")
	if tag == nil {
		tag = valueOfField(ctx, "Tag")
	}

	if page == "1" {
		url = "/"
	} else {
		url = fmt.Sprintf(`/page/%s/`, page)
	}
	if raymond.IsTrue(tag) && len(tag.(*model.Tag).Name) > 0  {
		tagName := tag.(*model.Tag).Name
		url = fmt.Sprintf("/tag/%s%s", tagName, url)
	}

	return joinPath(blog.URLPrefix, url)
}
