package controller

import (
	"github.com/beeender/dailyiker/model"
	"github.com/beeender/dailyiker/view"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"net/url"
	"path/filepath"
)

type Blog struct {
	Config
	Echo        *echo.Echo
	HostAndPort string
	RootURL     *url.URL
	ContentDir  string

	site  model.Site
	query model.Query

}

type Config struct {
	PostsPerPage int
}

func (blog *Blog) Start() error {
	// Init database
	db, err := gorm.Open("sqlite3", blog.dbPath())
	if err != nil {
		return err
	}
	db.LogMode(true)
	db.SetLogger(blog.Echo.Logger)
	blog.query = &model.DBDataQuery{DB: db}
	defer db.Close()

	// Load basic site information
	blog.site.Load(blog.query)

	// Load theme
	ren, err := view.NewTheme(blog.themePath())
	if err != nil {
		return err
	}
	blog.Echo.Renderer = ren

	// Register routes
	if err := blog.initRoute(); err != nil {
		return err
	}
	return blog.Echo.Start(blog.HostAndPort)
}

func (blog *Blog) dbPath() string {
	return filepath.Join(blog.ContentDir, "data/dailyiker.db")
}

func (blog *Blog) themePath() string {
	return filepath.Join(blog.ContentDir, "themes/"+blog.site.Theme)
}
