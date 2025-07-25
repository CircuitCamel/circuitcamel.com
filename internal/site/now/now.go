package now

import (
	"net/http"
	"os"
	"text/template"

	"circuitcamel.com/internal/models"
	"circuitcamel.com/internal/utils"
)

func Now(w http.ResponseWriter, r *http.Request) {
	var index models.Page

	databytes, err := os.ReadFile("content/now.md")
	if err != nil {
		utils.ErrPage(w, r, 500)
	}

	index.Body = utils.MdToHTML(databytes)

	tmpl, err := template.ParseFiles(
		"static/templates/head.html",
		"static/templates/footer.html",
	)

	tmpl.ExecuteTemplate(w, "base", index)
}
