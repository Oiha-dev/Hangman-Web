package back

import (
	"fmt"
	"net/http"
)

func StartServer() {
	fmt.Println("(http://localhost:8080) - Server started on port 8080")
	fmt.Println("Press Ctrl+C to stop the server")

	http.HandleFunc("/", Index)
	http.HandleFunc("/submit", LoginSubmit)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("internal/web/front"))))
	http.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.Dir("data"))))

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
