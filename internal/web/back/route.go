package back

import (
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "login/index")
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("internal/web/front/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
