package model

import (
	"github.com/aymerick/raymond"
	"time"
)

type Post struct {
	ID string
	UUID string
	Title        string
	PlainText string `gorm:"column:plaintext"`
	HTML string
	FeatureImage string `gorm:"column:feature_image" handlebars:"feature_image"`
	CustomExcerpt string `gorm:"column:custom_excerpt"`

	PublishedAt time.Time `gorm:"column:published_at"`

	Author User	`gorm:"-"`
	Authors Users	`gorm:"-"`
	AuthorID string `gorm:"column:author_id"`

	URL string `gorm:"-" handlebars:"url"`
	Content raymond.SafeString `gorm:"-"`
}

type PostsQuery interface {
	RecentPosts(count int) []Post
	PostByTitle(title string) *Post
}

// Set User's table name to be `profiles`
func (Post) TableName() string {
	return "posts"
}

func (q *DBDataQuery) RecentPosts(count int) []Post {
	var posts []Post
	q.DB.Where("status = 'published'").Order("published_at desc").Limit(count).Find(&posts)
	for i := range posts {
		q.queryPostAuthors(&posts[i])
	}
	return posts
}

func (q *DBDataQuery) PostByTitle(title string) *Post {
	var post Post
	if q.DB.Where("title = ?", title).First(&post).RecordNotFound() {
		return nil
	}
	q.queryPostAuthors(&post)
	return &post
}

func (q *DBDataQuery) queryPostAuthors(post *Post)  {
	q.DB.Where("id = ?", post.AuthorID).First(&post.Author)
	q.DB.Table("users").
		Joins("JOIN posts_authors ON posts_authors.author_id = users.id").
		Where("posts_authors.post_id = ?", post.ID).
		Find(&post.Authors)
}
