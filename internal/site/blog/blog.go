package blog

import (
	"net/http"

	"circuitcamel.com/internal/cache"
	"circuitcamel.com/internal/models"
	"circuitcamel.com/internal/utils"
	"github.com/gorilla/mux"
)

func BlogList(w http.ResponseWriter, r *http.Request) {
	cache.BlogListTmpl.ExecuteTemplate(w, "bloglist", cache.BlogPosts)
}

func Blog(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	var selected models.BlogPost
	found := false
	for _, b := range cache.BlogPosts {
		if b.Slug == slug {
			selected = b
			found = true
			break
		}
	}

	if !found {
		utils.ErrPage(w, r, 404)
		return
	}

	cache.PageTmpl.ExecuteTemplate(w, "page", selected)
}

func Latest(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/blog/"+cache.BlogPosts[0].Slug, http.StatusSeeOther)
}
