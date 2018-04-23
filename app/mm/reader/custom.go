package reader

import (
	"fmt"
	"net/http"
)

// CustomHandler handler
func CustomHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler Custom request")
	w.Write([]byte("Custom"))
}
