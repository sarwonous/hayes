package reader

import (
	"fmt"
	"net/http"
)

// PDPHandler page
func PDPHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handle PDP")
	w.Write([]byte("a"))
}
