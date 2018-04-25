package blog

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/unicolony/hayes/app/blog/reader"
)

// App App
type App struct {
}

// Init App
func (m *App) Init() {
	app := mux.NewRouter()
	app.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("public"))))
	app.HandleFunc("/{id:[a-z0-9-]+}.html", reader.Single)
	log.Fatal(http.ListenAndServe(":8000", app))
}

// Name App
func (m *App) Name() string {
	return "blog"
}

// Command app
func (m *App) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "blog",
		Short: "blog",
		Long:  "blog",
		Run: func(c *cobra.Command, args []string) {
			m.Init()
		},
	}
	return cmd
}

// NewApp App
func NewApp() *App {
	return &App{}
}
