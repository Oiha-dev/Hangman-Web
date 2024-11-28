package back

import (
	"fmt"
	classic_utils "hangman-web/pkg/hangman-classic/pkg/utils"
	"hangman-web/pkg/utils"
	"html/template"
	"net/http"
	"sync"
)

var (
	multiPageUserCount int
	userCountMutex     sync.Mutex
)

func multi(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"split":    utils.Split,
		"contains": classic_utils.ContainsStr,
	}

	tmpl, err := template.New("index.gohtml").Funcs(funcMap).ParseFiles("internal/web/front/multi/index.gohtml")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}

	userCountMutex.Lock()
	data := struct {
		UserCount int
	}{
		UserCount: multiPageUserCount,
	}
	userCountMutex.Unlock()

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

func multiServerSentEvents(w http.ResponseWriter, r *http.Request) {
	// Increment user count
	userCountMutex.Lock()
	multiPageUserCount++
	fmt.Println("User count:", multiPageUserCount)
	userCountMutex.Unlock()

	// Keep the connection open
	<-r.Context().Done()

	// Decrement user count when connection closes
	userCountMutex.Lock()
	multiPageUserCount--
	fmt.Println("User count:", multiPageUserCount)
	userCountMutex.Unlock()
}
