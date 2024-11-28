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
	listNames          []string
)

func waiting(w http.ResponseWriter, r *http.Request) {
	/*
		This function is used to display the waiting page
		by getting the name of the player
	*/
	name := r.FormValue("name")
	http.SetCookie(w, &http.Cookie{
		Name:    "playerName",
		Value:   name,
		Path:    "/",
		Expires: time.Now().Add(24 * time.Hour),
	})
	if multiPageUserCount >= 2 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	funcMap := template.FuncMap{
		"split":    utils.Split,
		"contains": classic_utils.ContainsStr,
	}

	tmpl, err := template.New("index.gohtml").Funcs(funcMap).ParseFiles("internal/web/front/waiting/index.gohtml")
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
func waitingSSE(w http.ResponseWriter, r *http.Request) {
	/*
		This function is used to stream updates for the user count
		using Server-Sent Events (SSE)
	*/

	// Set headers for SSE
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	userCountMutex.Lock()
	multiPageUserCount++
	fmt.Println("User count:", multiPageUserCount)
	userCountMutex.Unlock()

	for {
		select {
		case <-r.Context().Done():
			userCountMutex.Lock()
			multiPageUserCount--
			fmt.Println("User count:", multiPageUserCount)
			userCountMutex.Unlock()
			return
		default:
			userCountMutex.Lock()
			if multiPageUserCount >= 2 {
				fmt.Fprintf(w, "event: redirect\ndata: /multi\n\n")
				w.(http.Flusher).Flush()
				userCountMutex.Unlock()
				return
			}
			fmt.Println("User count:", multiPageUserCount)
			w.(http.Flusher).Flush()
			userCountMutex.Unlock()
			time.Sleep(1 * time.Second)
		}
	}
}
