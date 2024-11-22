package back

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func StartServer() {
	fmt.Println("(http://localhost:8080) - Server started on port", port[1:])
	fmt.Println("Press Ctrl+C to stop the server")

	http.HandleFunc("/", Index)
	http.HandleFunc("/submit", LoginSubmit)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("internal/web/front/login"))))

	err := http.ListenAndServe("127.0.0.1"+port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
