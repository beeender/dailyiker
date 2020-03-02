package view

import (
	"github.com/aymerick/raymond"
	"github.com/beeender/dailyiker/model"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"io"
)

const (
	RenderIndex = "index"
	RenderPost  = "post"
	RenderTags  = "tags"
	RenderTag  = "tag"
)

func (theme *Theme) Render(w io.Writer, name string, args interface{}, c echo.Context) error {
	switch name {
	case RenderIndex:
		return theme.renderIndex(w, args, c)
	case RenderPost:
		return theme.renderPost(w, args, c)
	case RenderTags:
		return theme.renderTags(w, args, c)
	case RenderTag:
		return theme.renderTag(w, args, c)
	}
	return echo.ErrNotFound
}

func (theme *Theme) RenderPagination(pagination model.Pagination) raymond.SafeString {
	template := theme.partials["pagination"]
	if template == nil {
		return ""
	}
	ret, _ := template.Exec(pagination)
	return raymond.SafeString(ret)
}

func (theme *Theme) RenderNavigation(ctx map[string]interface{}) raymond.SafeString {
	template := theme.partials["navigation"]
	if template == nil {
		return ""
	}
	ret, err := template.Exec(ctx)
	if err != nil {
		return ""
	}
	return raymond.SafeString(ret)
}

func (theme *Theme) renderTags(w io.Writer, args interface{}, _ echo.Context) error {
	ctx := args.(map[string]interface{})
	template := theme.bases["custom-tag-archive"]
	if template == nil {
		return echo.ErrNotFound
	}
	body, err := template.Exec(ctx)
	if err != nil {
		return err
	}
	ctx["body"] = body
	result, err := theme.bases["default"].Exec(ctx)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(result))
	return err
}

func (theme *Theme) renderIndex(w io.Writer, args interface{}, _ echo.Context) error {
	ctx := args.(map[string]interface{})
	body, err := theme.bases["index"].Exec(args)
	ctx["body"] = body

	result, err := theme.bases["default"].Exec(ctx)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(result))
	return err
}

func (theme *Theme) renderTag(w io.Writer, args interface{}, _ echo.Context) error {
	ctx := args.(map[string]interface{})
	body, err := theme.bases["tag"].Exec(args)
	ctx["body"] = body

	result, err := theme.bases["default"].Exec(ctx)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(result))
	return err
}

func (theme *Theme) renderPost(w io.Writer, args interface{}, _ echo.Context) error {
	ctx := args.(map[string]interface{})
	body, err := theme.bases["post"].Exec(args)
	ctx["body"] = body

	result, err := theme.bases["default"].Exec(ctx)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(result))
	return err
}
