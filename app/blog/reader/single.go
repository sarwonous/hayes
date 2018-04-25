package reader

import (
	"net/http"
)

// Single Articles
func Single(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Yihayyy..."))
}
