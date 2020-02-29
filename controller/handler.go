package controller

import (
	"github.com/aymerick/raymond"
	"github.com/beeender/dailyiker/model"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"strings"
)

func (blog *Blog) indexHandler(c echo.Context) error {
	posts := blog.query.PostsAtPage(blog.Config.PostsPerPage, 1)
	for i := range posts {
		blog.completePostInfo(&posts[i])
	}
	args := blog.metaArgs()
	args["posts"] = posts
	pagination := blog.createPagination(1)

	args["pagination"] = pagination
	return c.Render(http.StatusOK, "index", args)
}

func (blog *Blog) pageHandler(c echo.Context) error {
	params := c.ParamValues()
	if len(params) != 1 {
		return echo.ErrNotFound
	}
	pageStr := strings.Trim(params[0], "/")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return echo.ErrNotFound
	}
	pagination := blog.createPagination(page)

	posts := blog.query.PostsAtPage(blog.Config.PostsPerPage, page)
	for i := range posts {
		blog.completePostInfo(&posts[i])
	}

	args := blog.metaArgs()
	args["posts"] = posts
	args["pagination"] = pagination

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

	y, _ := strconv.Atoi(strList[0])
	m, _ := strconv.Atoi(strList[1])
	d, _ := strconv.Atoi(strList[2])
	t := strList[3]

	post := blog.query.PostByTitle(t)
	if post == nil {
		return echo.ErrNotFound
	}

	if y != post.PublishedAt.Year() || m != int(post.PublishedAt.Month()) || d != post.PublishedAt.Day() {
		return echo.ErrNotFound
	}

	post.Content = raymond.SafeString(post.HTML)
	args := blog.metaArgs()
	args["post"] = post
	return c.Render(http.StatusAccepted, "post", args)
}

/*
func (blog *Blog) apiV2ContentPostsHandler(c echo.Context) error {
	key := c.QueryParam("key")
	limit := c.QueryParam("limit")
	filter := c.QueryParam("filter")
	include := c.QueryParam("include")
	c.JSON(http.StatusOK, key + limit + filter +include)
	return nil
}
 */

func (blog *Blog) createPagination(page int) model.Pagination {
	pagination := blog.query.NewPagination("", blog.Config.PostsPerPage)
	pagination.Page = page
	next := pagination.Page + 1
	prev := pagination.Page - 1
	if next > pagination.Pages {
		next = 0
	}
	if prev < 1 {
		prev = 0
	}
	pagination.Next = next
	pagination.Prev = prev
	return pagination
}

func (blog *Blog) metaArgs() map[string]interface{} {
	args := map[string]interface{}{
		"site":         blog.site,
		"published_at": blog.query.LastUpdatedAt(),
		"meta_title":   blog.site.MetaTitle,
	}
	return args
}
