package index

import (
	"net/http"
	"os"

	"circuitcamel.com/internal/cache"
	"circuitcamel.com/internal/models"
	"circuitcamel.com/internal/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var index models.Page

	databytes, err := os.ReadFile("content/index.md")
	if err != nil {
		utils.ErrPage(w, r, 500)
	}

	index.Body = utils.MdToHTML(databytes)

	cache.PageTmpl.ExecuteTemplate(w, "page", index)
}
