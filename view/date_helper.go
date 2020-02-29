package view

import (
	"github.com/aymerick/raymond"
	"time"
)

func (theme *Theme) dateHelper(opts *raymond.Options) interface{} {
	var date time.Time
	var format string
	format = opts.HashStr("format")

	// TODO: Handle the hash param "date"

	ctx := opts.Ctx()
	v := valueOfField(ctx, "PublishedAt")
	if v == nil {
		v = valueOfMap(ctx, "published_at")
	}
	if v == nil {
		return nil
	}

	date = v.(time.Time)
	// TODO: Hardcoded the time format. There are some works need to support
	// all formats in Ghost/core/frontend/helpers/date.js
	if format == "YYYY" {
		format = "2006"
	} else {
		format = "2006-01-02"
	}
	return date.Format(format)
}

