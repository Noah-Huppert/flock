// Main HTTP server package for Squad Up.
package main

// Import deps.
import (
	// Golang deps.
	"fmt"
	"net/http"
)

// Main entry point of program.
func main() {
	// New HTTP router.
	mux := http.NewServeMux()

	// Serves index.html at root.
	mux.HandleFunc("/", func (w http.ResponseWriter, r*http.Request) {
		// File path (3rd arg) is relative to Go cli wkr dir (go/src/squad-up)
		http.ServeFile(w, r, "client/views/index.html")
	})

	// Start listening on any host, port 5000.
	err := http.ListenAndServe(":5000", mux)
	if err != nil { // If err print to console.
		fmt.Println("Error starting HTTP server on :5000 -> {}", err)
	} else { // Else print normal OK status message.
		fmt.Println("Listening on :5000")
	}
}
