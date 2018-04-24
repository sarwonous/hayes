package reader

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unicolony/hayes/service/meta"
)

// BrandHandler handler
func BrandHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println("handler Brand request")
	brandMeta := meta.GetMetaByURL(fmt.Sprintf("/brand/%s/%s", vars["id"], vars["name"]))
	// jsonResponse, err := json.Marshal(brandMeta)
	// if err != nil {
	// 	w.Write([]byte("Error parsing"))
	// }
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	tmpl.Execute(w, brandMeta)
}
