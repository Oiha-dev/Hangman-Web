package back

import (
	"fmt"
	classic_utils "hangman-web/pkg/hangman-classic/pkg/utils"
	"hangman-web/pkg/utils"
	"html/template"
	"net/http"
	"sync"
	"time"
)

var (
	multiPageUserCount int
	userCountMutex     sync.Mutex
)

func multi(w http.ResponseWriter, r *http.Request) {
	if multiPageUserCount >= 2 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
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

// Stream updates for user count using SSE
func multiServerSentEvents(w http.ResponseWriter, r *http.Request) {
	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Increment user count
	userCountMutex.Lock()
	multiPageUserCount++
	fmt.Println("User count:", multiPageUserCount)
	userCountMutex.Unlock()

	// Send initial user count
	fmt.Fprintf(w, "data: %d\n\n", multiPageUserCount)
	w.(http.Flusher).Flush()

	// Keep the connection open and send updates
	for {
		select {
		case <-r.Context().Done():
			// Decrement user count when connection closes
			userCountMutex.Lock()
			multiPageUserCount--
			fmt.Println("User count:", multiPageUserCount)
			userCountMutex.Unlock()
			return
		default:
			// Send updated user count
			userCountMutex.Lock()
			fmt.Fprintf(w, "data: %d\n\n", multiPageUserCount)
			userCountMutex.Unlock()
			w.(http.Flusher).Flush()
			time.Sleep(1 * time.Second) // Adjust the interval as needed
		}
	}
}
