package back

import (
	"net/http"
)

func history(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "history/index", nil)
}
