package reader

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// Tempalte File
var (
	TemplateFile = ""
)

// SetupRouter s
func SetupRouter(m *mux.Router, routes []string) {
	source := viper.GetString("source")
	TemplateFile = fmt.Sprintf("%s/%s", source, "index.html")
	m.StrictSlash(true)
	m.HandleFunc("/{name:[a-z0-9]+}-{id:[0-9]+}.html", PDPHandler).Methods("GET")
	m.HandleFunc("/p-{id:[0-9]+}/{name:[a-z0-9]+}", PCPHandler).Methods("GET")
	m.HandleFunc("/brand/{id:[0-9]+}/{name:[a-z0-9]+}", BrandHandler).Methods("GET")
	m.HandleFunc("/store/{id:[0-9]+}/{name:[a-z0-9]+}", StoreHandler).Methods("GET")
	m.HandleFunc("/mau-gaya-itu-gampang/{campaign_name}-{campaign_id}/{post_id}/{icode}", CampaignHandler).Methods("GET")
	m.HandleFunc("/", HomeHandler).Methods("GET")
	for _, name := range routes {
		m.HandleFunc(name, CustomHandler).Methods("GET")
	}
	m.PathPrefix("/assets/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(source))))
	m.PathPrefix("/{name}.js").Handler(http.StripPrefix("/", http.FileServer(http.Dir(source))))
	m.PathPrefix("/{name}.css").Handler(http.StripPrefix("/", http.FileServer(http.Dir(source))))
	m.PathPrefix("/robots.txt").Handler(http.StripPrefix("/", http.FileServer(http.Dir(source))))
	m.NotFoundHandler = http.HandlerFunc(NotFoundHandler)
}
