package routes

import (
	"net/http"
	"path/filepath"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// Define the path to the static directory
	staticDir := "static"

	// Serve the index.html file
	http.ServeFile(w, r, filepath.Join(staticDir, "index.html"))
}
