package now

import (
	"net/http"
	"os"

	"circuitcamel.com/internal/cache"
	"circuitcamel.com/internal/models"
	"circuitcamel.com/internal/utils"
)

func Now(w http.ResponseWriter, r *http.Request) {
	var now models.Page

	databytes, err := os.ReadFile("content/now.md")
	if err != nil {
		utils.ErrPage(w, r, 500)
	}

	now.Body = utils.MdToHTML(databytes)

	cache.PageTmpl.ExecuteTemplate(w, "page", now)
}
