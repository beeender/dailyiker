package controller

import (
	"github.com/aymerick/raymond"
	"github.com/beeender/dailyiker/model"
)

func (blog *Blog) navigationHelper(opts *raymond.Options) interface{} {
	blogCtx := opts.Ctx()
	site := valueOfMap(blogCtx, "site").(model.Site)
	ctx := map[string]interface{} {
		"navigation": site.Navigation,
	}

	return blog.theme.RenderNavigation(ctx)
}
