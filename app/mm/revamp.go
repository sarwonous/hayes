package mm

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"./reader"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

// App MMApp
type App struct {
	router *mux.Router
}

func (m *App) handler() {
	// pdp
	m.router.HandleFunc("/{name:[a-z0-9]+}-{id:[0-9]+}.html", reader.PDPHandler)
	// pcp
	m.router.HandleFunc("/p-{id:[0-9]+}/{name:[a-z0-9]+}", reader.PCPHandler)
	// brand
	m.router.HandleFunc("/brand/{id:[0-9]+}/{name:[a-z0-9]+}", reader.BrandHandler)
	// home
	m.router.HandleFunc("/", reader.HomeHandler)
	// custom url
	m.router.PathPrefix("/assets/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./html/"))))
	m.router.PathPrefix("/{name}.js").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./html/"))))
	m.router.PathPrefix("/{name}.css").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./html/"))))
	m.router.PathPrefix("/").HandlerFunc(reader.CustomHandler)
}

// RevampInit RevampInit init
func (m *App) RevampInit() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	m.router = mux.NewRouter()
	m.handler()
	srv := &http.Server{
		Addr: "0.0.0.0:8080",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      m.router, // Pass our instance of gorilla/mux in.
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Shutting down")
	os.Exit(0)
}

// Init Init
func (m *App) Init() {
	fmt.Println("App inited")
	m.RevampInit()
}

// Name Name
func (m *App) Name() string {
	return "mm"
}

// Command Command
func (m *App) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mm",
		Short: "mm",
		Long:  "mm",
		Run: func(c *cobra.Command, args []string) {
			m.Init()
		},
	}
	return cmd
}

// NewMM NewMM
func NewMM() *App {
	app := &App{}
	return app
}
