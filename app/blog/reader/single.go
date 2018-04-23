package reader

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Single Articles
func Single(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)
}

// Category Pages
func Category(w http.ResponseWriter, r *http.Request) {

}
