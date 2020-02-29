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
	return tags
}
