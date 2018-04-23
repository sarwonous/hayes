package blog

import (
	"log"
	"net/http"

	"./reader"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

// App App
type App struct {
}

// Init App
func (m *App) Init() {
	app := mux.NewRouter()

	app.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("public"))))

	// app.HandleFunc("/", cont.MHome)

	// app.HandleFunc("/{cat1:[a-z0-9-]+}/{cat2:[a-z0-9-]+}/{cat3:[a-z0-9-]+}/{id:[a-z0-9-]+}.html", cont.MRead)
	// app.HandleFunc("/{cat1:[a-z0-9-]+}/{cat2:[a-z0-9-]+}/{id:[a-z0-9-]+}.html", cont.MRead)
	// app.HandleFunc("/{cat1:[a-z0-9-]+}/{id:[a-z0-9-]+}.html", cont.MRead)
	// app.HandleFunc("/tags/{tags:[a-z0-9]+}", cont.MTags)
	// app.HandleFunc("/{category:[a-z0-9]+}", cont.MCategories)
	app.HandleFunc("/{id:[a-z0-9-]+}.html", reader.Single)
	// app.Start(config.GetString("host"))
	// http.Handle("/", app)/
	log.Fatal(http.ListenAndServe(":8000", app))
}

// Name App
func (m *App) Name() string {
	return "blog"
}

func (m *App) Command() *cobra.Command {
	cmd := oist===oist==
}

// NewApp App
func NewApp() {
	return &App{}
}
