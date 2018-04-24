package reader

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cast"
	"github.com/unicolony/hayes/service/meta"
)

// PCPHandler handler
func PCPHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	categoryMeta := meta.GetMetaCategory(cast.ToInt(vars["id"]))
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	tmpl.Execute(w, categoryMeta)
}
