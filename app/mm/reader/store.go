package reader

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unicolony/hayes/service/meta"
)

// StoreHandler handler
func StoreHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	brandMeta := meta.GetMetaByURL(fmt.Sprintf("/store/%s/%s", vars["id"], vars["name"]))
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	tmpl.Execute(w, brandMeta)
}
