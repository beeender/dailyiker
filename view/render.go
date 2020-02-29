package view

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"io"
)

const (
	RenderIndex = "index"
	RenderPost  = "post"
)

func (theme *Theme) Render(w io.Writer, name string, args interface{}, c echo.Context) error {
	switch name {
	case RenderIndex:
		return theme.renderIndex(w, args, c)
	case RenderPost:
		return theme.renderPost(w, args, c)
	}
	return echo.ErrNotFound
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
