package reader

import (
	"fmt"
	"net/http"
)

// BrandHandler handler
func BrandHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler Brand request")
	w.Write([]byte("Brand"))
}
