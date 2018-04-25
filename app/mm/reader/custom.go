package reader

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/unicolony/hayes/service/meta"
)

// CustomHandler handler
func CustomHandler(w http.ResponseWriter, r *http.Request) {
	customMeta := meta.GetMetaByURL(fmt.Sprintf("%s", r.URL.Path))
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	tmpl.Execute(w, customMeta)
}
