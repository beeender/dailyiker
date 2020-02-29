package model

type Site struct {
	Title       string
	MetaTitle   string `handlebars:"meta_title"`
	Description string
	CoverImage  string `handlebars:"cover_image"`
	Logo        string
	Theme       string
	URL         string
}

func (site *Site) Load(q SettingsQuery) {
	if s := q.SettingByKey("title"); s != nil {
		site.Title = s.Value
	}
	if s := q.SettingByKey("meta_title"); s != nil {
		site.MetaTitle = s.Value
	}
	if s := q.SettingByKey("description"); s != nil {
		site.Description = s.Value
	}
	if s := q.SettingByKey("cover_image"); s != nil {
		site.CoverImage = s.Value
	}
	if s := q.SettingByKey("logo"); s != nil {
		site.Logo = s.Value
	}
	site.Theme = "zvikov"
}
