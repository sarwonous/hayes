package reader

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unicolony/hayes/service/meta"
)

// CustomHandler handler
func CustomHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	brandMeta := meta.GetMetaByURL(fmt.Sprintf("/%s", vars["name"]))
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	tmpl.Execute(w, brandMeta)
}
