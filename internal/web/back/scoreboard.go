package back

import (
	"fmt"
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Hangman Game!")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
}

func scoreboard(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "scoreboard")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("internal/web/front/" + tmpl + "/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}

func RouteScoreboard() {
	http.HandleFunc("/", home)
	http.HandleFunc("/scoreboard", scoreboard)

	// Serve static files
	fs := http.FileServer(http.Dir("internal/web/front/scoreboard"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Serve font and image files
	dataFs := http.FileServer(http.Dir("data"))
	http.Handle("/data/", http.StripPrefix("/data/", dataFs))

	fmt.Println("Server is running on port 8080")
	fmt.Println("(http://localhost:8080/scoreboard) to view the scoreboard")

	http.ListenAndServe(":8080", nil)
}
