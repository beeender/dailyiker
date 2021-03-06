package model

import (
	"fmt"
	"time"
)

type Tag struct {
	ID        string
	Name      string
	Slug      string
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string

	URL string `gorm:"-" handlebars:"url"`
}

type TagsQuery interface {
	Tags(count int) []Tag
	TagByName(name string) *Tag
	TagsByPost(post *Post) []Tag
}

func (Tag) Table() string {
	return "tags"
}

func (tag Tag) String() string {
	return tag.Name
}

func (q *DBDataQuery) Tags(count int) []Tag {
	var tags []Tag
	db := q.DB
	limit := ""
	if count > 0 {
		limit = fmt.Sprintf("LIMIT %d", count)
	}
	// Sort the tags by the occurrences
	sql := fmt.Sprintf(`
SELECT tags.*, COUNT(pt.tag_id) count
FROM posts_tags pt, tags
WHERE tags.id = pt.tag_id
GROUP BY tag_id
ORDER BY count DESC
%s
`, limit)

	db = db.Raw(sql).Scan(&tags)
	for i := range tags {
		tags[i].makeURL()
	}
	return tags
}

func (q *DBDataQuery) TagByName(name string) *Tag {
	var tag Tag
	if q.DB.Where("name = ?", name).Find(&tag).RecordNotFound() {
		return nil
	}
	tag.makeURL()
	return &tag
}

func (q *DBDataQuery) TagsByPost(post *Post) []Tag {
	var tags []Tag
	q.DB.Table("tags").
		Joins("JOIN posts_tags ON tags.id = posts_tags.tag_id AND posts_tags.post_id = ?", post.ID).
		Find(&tags)
	for i := range tags {
		tags[i].makeURL()
	}
	return tags
}

func (tag *Tag)makeURL() {
	tag.URL = fmt.Sprintf("/tag/%s/", tag.Name)
}
