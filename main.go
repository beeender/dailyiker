package main

import (
	"github.com/beeender/dailyiker/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"net/url"
	"path"
)

func main() {
	b := controller.Blog{
		Echo:        echo.New(),
		HostAndPort: "127.0.0.1:8084",
		ContentDir:  "/home/cc/repo/dailyiker/content",
	}

	b.Echo.Logger.SetLevel(log.DEBUG)
	b.Echo.Use(middleware.Logger())
	b.Echo.Use(middleware.Recover())

	b.Config.PostsPerPage = 5
	b.Config.URLPrefix = "/blog/"
	b.Config.RootURL = "https://www.ikeriker.org"

	root := path.Join(b.Config.RootURL, b.Config.URLPrefix)
	u, err := url.Parse(root)
	if err != nil {
		b.Echo.Logger.Fatal(err)
	}
	b.RootURL = *u

	b.Echo.Logger.Fatal(b.Start())
}
