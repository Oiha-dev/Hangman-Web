package back

import (
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "login/index", nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	/*
		This function is used to render the templates using the data sent
	*/
	t, err := template.ParseFiles("internal/web/front/" + tmpl + ".gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data)
}
