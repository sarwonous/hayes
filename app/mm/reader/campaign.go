package reader

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unicolony/hayes/service/meta"
)

// CampaignHandler handler
func CampaignHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	meta := meta.GetMetaByURL(fmt.Sprintf("/mau-gaya-itu-gampang/%s-%s/%s/%s", vars["campaign_name"], vars["campaign_id"], vars["post_id"], vars["icode"]))
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	tmpl.Execute(w, meta)
}
