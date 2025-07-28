package teapot

import (
	"net/http"

	"circuitcamel.com/internal/utils"
)

func Teapot(w http.ResponseWriter, r *http.Request) {
	utils.ErrPage(w, r, 418)
}
