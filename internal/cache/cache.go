package cache

import (
	"text/template"
	"time"

	"circuitcamel.com/internal/models"
)

var (
	BlogPosts    []models.BlogPost
	PageTmpl     *template.Template
	BlogListTmpl *template.Template
)

func LoadAll() {
	var err error

	for {
		BlogPosts, err = getBlogPosts()
		if err != nil {
			return
		}
		err = loadTemplates()
		if err != nil {
			return
		}
		time.Sleep(time.Hour)
	}
}
