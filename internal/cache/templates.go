package cache

import "text/template"

func loadTemplates() error {
	var err error

	PageTmpl, err = template.ParseFiles(
		"static/templates/page.html",
		"static/templates/head.html",
		"static/templates/footer.html",
	)
	if err != nil {
		return err
	}

	BlogListTmpl, err = template.ParseFiles(
		"static/templates/bloglist.html",
		"static/templates/head.html",
		"static/templates/footer.html",
	)

	return err
}
