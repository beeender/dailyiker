package controller

import (
"github.com/aymerick/raymond"
	"github.com/beeender/dailyiker/model"
)

func (blog *Blog) paginationHelper(opts *raymond.Options) interface{} {
	ctx := opts.Ctx().(map[string]interface{})
	pagination := ctx["pagination"].(model.Pagination)

	return blog.theme.RenderPagination(pagination)
}
