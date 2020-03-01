package model

type Navigation struct {
	Label string
	URL string `gorm:"-" handlebars:"url"`
	Current bool
	Slug string
}

func NewNavigations() []Navigation {
	navs := []Navigation {
		{
			Label:`Home`,
			URL: `/`,
		},
		{
			Label:`Topics`,
			URL: `/tag/`,
		},
	}
	return navs
}
