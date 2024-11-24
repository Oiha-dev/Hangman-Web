package back

import (
	"net/http"
	"time"
)

func LoginSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		http.SetCookie(w, &http.Cookie{
			Name:    "playerName",
			Value:   name,
			Path:    "/",
			Expires: time.Now().Add(24 * time.Hour),
		})
		http.Redirect(w, r, "/game", http.StatusSeeOther)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
