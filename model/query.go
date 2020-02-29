package model

import (
	"github.com/jinzhu/gorm"
)

type Query interface {
	SettingsQuery
	PostsQuery
	TagsQuery
	PaginationQuery
}

type DBDataQuery struct {
	DB *gorm.DB
}
