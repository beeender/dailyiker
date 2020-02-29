package main

import (
	"github.com/beeender/dailyiker/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"net/url"
)

func main() {
	b := controller.Blog{
		Echo:        echo.New(),
		HostAndPort: "127.0.0.1:8084",
		ContentDir:  "/home/cc/repo/dailyiker/content",
	}
	b.RootURL, _ = url.Parse("http://127.0.0.1:8084")

	b.Config.PostsPerPage = 5

	b.Echo.Logger.SetLevel(log.DEBUG)
	b.Echo.Use(middleware.Logger())
	b.Echo.Use(middleware.Recover())
	b.Echo.Logger.Fatal(b.Start())
}
