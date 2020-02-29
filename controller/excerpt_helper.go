package controller

import (
	"github.com/aymerick/raymond"
	"github.com/beeender/dailyiker/model"
	"strconv"
	"unicode/utf8"
)

func (blog *Blog) excerptHelper(options *raymond.Options) interface{} {
	ctx := options.Ctx()
	post := ctx.(model.Post)
	excerpt := ""
	if len(post.CustomExcerpt) > 0 {
		excerpt = post.CustomExcerpt
	} else if len(post.PlainText) > 0 {
		excerpt = post.PlainText
	}
	words := 0
	words, _ = strconv.Atoi(options.HashStr("words"))
	if words == 0 {
		words = 40
	}
	excerpt = generateExcerpt(excerpt, words)
	return excerpt
}

func generateExcerpt(content string, words int) string {
	runeStr := []rune(content)
	// For Chinese, roughly take the first 3 * words characters
	end := utf8.RuneCountInString(content)
	l := words * 3
	if l < end {
		end = l
	}
	return string(runeStr[0:end])
}
