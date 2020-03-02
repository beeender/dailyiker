package model

import "strconv"

type Pagination struct {
	// page in set to display
	Page int
	// no. results per page, or 'all'
	Limit string
	// total no. pages in the full set
	Pages int
	// total no. items in the full set
	Total int
	// next page
	Next int
	// previous page
	Prev int
	Tag  *Tag
}

type PaginationQuery interface {
	NewPagination(tag string, postsPerPage int) *Pagination
}

func (q *DBDataQuery) NewPagination(tagName string, postsPerPage int) *Pagination {
	var pagination Pagination
	var tag *Tag
	if len(tagName) > 0 {
		tag = q.TagByName(tagName)
		if tag == nil {
			return nil
		}
	}

	db := q.DB.Table("posts").
		Where("status = 'published'")

	if tag != nil {
		pagination.Tag = tag
		db = db.Joins(
			"JOIN posts_tags ON posts.id = posts_tags.post_id AND posts_tags.tag_id = ? ",
			tag.ID)
	}

	db.Count(&pagination.Total)

	if postsPerPage > 0 {
		pagination.Limit = strconv.Itoa(postsPerPage)
		pagination.Pages = pagination.Total / postsPerPage
		if pagination.Total%postsPerPage > 0 {
			pagination.Pages += 1
		}
	} else {
		pagination.Limit = "all"
		pagination.Pages = 1
	}

	if pagination.Pages > 0 {
		pagination.Page = 1
	}

	return &pagination
}
