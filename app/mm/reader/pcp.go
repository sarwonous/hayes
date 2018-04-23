package reader

import (
	"fmt"
	"net/http"
)

// PCPHandler handler
func PCPHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handler pcp request")
	w.Write([]byte("PCP"))
}
