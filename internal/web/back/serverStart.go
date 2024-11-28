package back

import (
	"fmt"
	"net/http"
)

func StartServer() {
	/*
		This function is used to start the server
		and set the routes for the server
	*/
	fmt.Println("(http://localhost:8080) - Server started on port 8080")
	fmt.Println("Press Ctrl+C to stop the server")

	http.HandleFunc("/", index)
	http.HandleFunc("/submit", loginSubmit)
	http.HandleFunc("/scoreboard", scoreboard)
	http.HandleFunc("/history", history)
	http.HandleFunc("/game", gamePage)
	http.HandleFunc("/guess", guessLetter)
	http.HandleFunc("/fullword", fullWordGuess)
	http.HandleFunc("/end", endScreen)
	http.HandleFunc("/waiting", waiting)
	http.HandleFunc("/waiting/sse", waitingSSE)
	http.HandleFunc("/multi", multi)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("internal/web/front"))))
	http.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.Dir("data"))))

	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
