package reader

import (
	"html/template"
	"net/http"

	"github.com/unicolony/hayes/service/meta"
)

// PDPHandler page
func PDPHandler(w http.ResponseWriter, r *http.Request) {
	homeMeta := meta.GetMetaByURL("https://www.mataharimall.com")
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	tmpl.Execute(w, homeMeta)
}
