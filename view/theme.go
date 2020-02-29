package view

import (
	"github.com/aymerick/raymond"
	"os"
	"path/filepath"
	"strings"
)

type Theme struct {
	path       string
	assetsPath string
	partials   map[string]*raymond.Template
	bases      map[string]*raymond.Template
}

func NewTheme(path string) (*Theme, error) {
	theme := Theme{
		path:       path,
		assetsPath: filepath.Join(path, "assets"),
		partials:   map[string]*raymond.Template{},
		bases:      map[string]*raymond.Template{},
	}

	if err := theme.load(); err != nil {
		return nil, err
	}

	raymond.RegisterHelper("img_url", theme.imgURLHelper)
	raymond.RegisterHelper("excerpt", theme.excerptHelper)
	raymond.RegisterHelper("asset", theme.assetHelper)
	raymond.RegisterHelper("foreach", theme.foreachHelper)

	return &theme, nil
}

func (theme *Theme) load() error {
	walkFun := func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		if ext != ".hbs" {
			return nil
		}

		template, e := raymond.ParseFile(path)
		if e != nil {
			return e
		}

		strs := strings.SplitAfter(path, "partials")
		if len(strs) > 1 {
			name := strings.TrimSuffix(strs[1], ".hbs")
			name = strings.TrimPrefix(name, "/")
			theme.partials[name] = template
		} else {
			name := strings.TrimSuffix(filepath.Base(path), ".hbs")
			theme.bases[name] = template
		}
		return nil
	}

	if e := filepath.Walk(theme.path, walkFun); e != nil {
		return e
	}
	for k, v := range theme.partials {
		raymond.RegisterPartialTemplate(k, v)
	}
	return nil
}
