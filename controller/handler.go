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
	posts := blog.query.PostsAtPage(nil, blog.Config.PostsPerPage, 1)
	for i := range posts {
		blog.completePostInfo(&posts[i])
	}
	args := blog.metaArgs()
	args["posts"] = posts
	pagination := blog.createPagination("", 1)

	args["pagination"] = *pagination
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
	pagination := blog.createPagination("", page)

	posts := blog.query.PostsAtPage(nil, blog.Config.PostsPerPage, page)
	for i := range posts {
		blog.completePostInfo(&posts[i])
	}

	args := blog.metaArgs()
	args["posts"] = posts
	args["pagination"] = *pagination

	return c.Render(http.StatusOK, "index", args)
}

func (blog *Blog) tagPageHandler(c echo.Context) error {
	var err error
	params := c.ParamValues()
	if len(params) != 1 {
		return echo.ErrNotFound
	}

	params = strings.Split(params[0], "/")
	if len(params) != 2 && len(params) != 4{
		return echo.ErrNotFound
	}

	page := 1
	if len(params) > 2 {
		pageStr := strings.Trim(params[2], "/")
		page, err = strconv.Atoi(pageStr)
		if err != nil {
			return echo.ErrNotFound
		}
	}

	tag := strings.Trim(params[0], "/")

	pagination := blog.createPagination(tag, page)
	if pagination == nil {
		return echo.ErrNotFound
	}

	posts := blog.query.PostsAtPage(pagination.Tag, blog.Config.PostsPerPage, page)
	for i := range posts {
		blog.completePostInfo(&posts[i])
	}

	args := blog.metaArgs()
	args["posts"] = posts
	args["pagination"] = *pagination
	args["tag"] = *pagination.Tag

	return c.Render(http.StatusOK, "tag", args)
}

func (blog *Blog) tagHandler(c echo.Context) error {
	args := blog.metaArgs()
	return c.Render(http.StatusOK, "tags", args)
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

func (blog *Blog) createPagination(tagName string, page int) *model.Pagination {
	pagination := blog.query.NewPagination(tagName, blog.Config.PostsPerPage)
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
		"site":         blog.site.Clone(),
		"published_at": blog.query.LastUpdatedAt(),
		"meta_title":   blog.site.MetaTitle,
	}
	return args
}
