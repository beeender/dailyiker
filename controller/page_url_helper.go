package controller

import (
	"fmt"
	"github.com/aymerick/raymond"
)

func (blog *Blog) pageURLHelper(page string, _ *raymond.Options) interface{} {
	if page == "1" {
		return "/"
	}
	return fmt.Sprintf(`/page/%s/`, page)
}
