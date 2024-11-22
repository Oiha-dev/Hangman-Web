package back

import (
	"net/http"
	"text/template"
)

func LoginSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		data := map[string]string{"Name": name}
		t, err := template.ParseFiles("internal/web/front/login/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		t.Execute(w, data)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
