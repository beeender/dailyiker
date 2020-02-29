package controller

import (
	"fmt"
	"github.com/beeender/dailyiker/model"
	"path"
	"path/filepath"
	"regexp"
)

func (blog *Blog) initRoute() error {
	blog.Echo.Static("/content/images", filepath.Join(blog.ContentDir, "images"))
	blog.Echo.Static("/assets", filepath.Join(blog.themePath(), "assets"))
	blog.Echo.GET("/", blog.indexHandler)
	blog.Echo.GET("/*/", blog.postHandler)

	return nil
}

func (blog *Blog) completePostInfo(post *model.Post) {
	if len(post.FeatureImage) == 0 {
		blog.completePostFeaturedImage(post)
	}
	if len(post.FeatureImage) > 0 {
		post.FeatureImage = path.Join(blog.RootURL.Path, post.FeatureImage)
	}
	blog.completePostUrl(post)
}

func (blog *Blog) completePostFeaturedImage(post *model.Post) {
	images := findImages(post.HTML)
	if len(images) > 0 {
		post.FeatureImage = images[0]
	}
}

func (blog *Blog) completePostUrl(post *model.Post) {
	t := post.PublishedAt
	url := fmt.Sprintf(`/%04d/%02d/%02d/%s/`,
		t.Year(), t.Month(), t.Day(),
		post.Title)
	post.URL = url
}

func (blog *Blog) completeTagsInf(tags []model.Tag) {
	for i := range tags {
		blog.completeTagInf(&tags[i])
	}
}

func (blog *Blog) completeTagInf(tag *model.Tag) {
	tag.URL = fmt.Sprintf(`/tag/%s/`, tag.Name)
}

var imgRE = regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)

// if your img's are properly formed with doublequotes then use this, it's more efficient.
// var imgRE = regexp.MustCompile(`<img[^>]+\bsrc="([^"]+)"`)
func findImages(htm string) []string {
	imgs := imgRE.FindAllStringSubmatch(htm, -1)
	out := make([]string, len(imgs))
	for i := range out {
		out[i] = imgs[i][1]
	}
	return out
}
