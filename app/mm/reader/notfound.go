package reader

import (
	"html/template"
	"net/http"

	"github.com/unicolony/hayes/service/meta"
)

// NotFoundHandler handler
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	meta := meta.GetMetaByURL("/not-found")
	tmpl := template.Must(template.ParseFiles(TemplateFile))
	tmpl.Execute(w, meta)
}
