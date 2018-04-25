package reader

import (
	"html/template"
	"net/http"

	"github.com/spf13/viper"
	"github.com/unicolony/hayes/service/meta"
)

// HomeHandler handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	homeMeta := meta.GetMetaByURL(viper.GetString("base_url"))
	tmpl := template.Must(template.ParseFiles(TemplateFile))
	tmpl.Execute(w, homeMeta)
}
