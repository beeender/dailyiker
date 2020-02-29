package controller

import (
	"github.com/aymerick/raymond"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"strings"
)

func (blog *Blog) indexHandler(c echo.Context) error {
	posts := blog.query.RecentPosts(blog.Config.PostsPerPage)
	for i := range posts {
		blog.completePostInfo(&posts[i])
	}
	args := map[string]interface{}{
		"site":  blog.site,
		"posts": posts,
	}
	return c.Render(http.StatusOK, "index", args)
}

func (blog *Blog) postHandler(c echo.Context) error {
	params := c.ParamValues()

	if len(params) != 1 {
		return echo.ErrNotFound
	}

	strList := strings.Split(params[0], "/")
	if len(strList) != 5 {
		return echo.ErrNotFound
	}

	y,_ := strconv.Atoi(strList[0])
	m,_ := strconv.Atoi(strList[1])
	d,_ := strconv.Atoi(strList[2])
	t := strList[3]

	post := blog.query.PostByTitle(t)
	if post == nil {
		return echo.ErrNotFound
	}

	if y != post.PublishedAt.Year() || m != int(post.PublishedAt.Month()) || d != post.PublishedAt.Day() {
		return echo.ErrNotFound
	}

	post.Content = raymond.SafeString(post.HTML)
	args := map[string]interface{}{
		"site":  blog.site,
		"post": post,
	}
	return c.Render(http.StatusAccepted, "post", args)
}

