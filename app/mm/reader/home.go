package reader

import (
	"fmt"
	"net/http"
)

// HomeHandler handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler Home request")
	w.Write([]byte("Home"))
}
