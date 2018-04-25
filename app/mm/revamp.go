package mm

import (
	"context"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/viper"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/unicolony/hayes/app/mm/reader"
)

// App MMApp
type App struct {
	router *mux.Router
}

var (
	Templates    map[string]*template.Template
	customRoutes = []string{
		"/",
		"/mau-gaya-itu-gampang",
		"/mau-gaya-itu-gampang/{campaign_name}-{campaign_id}/{post_id}/{icode}",
		"/search",
		"/products",
		"/searchnotfound",
		"/sub-category/",
		"/category/",
		"/subcategory/{categoryId}",
		"/brands",
		"/catalogcategory",
		"/product/comments/{id}",
		"/product/reviews/{id}",
		"/product/guide",
		"/login",
		"/logout",
		"/register",
		"/registered",
		"/forgot-password",
		"/user/newpassword",
		"/promo",
		"/promo/{type}",
		"/profile",
		"/profile-edit",
		"/lovelist",
		"/cart",
		"/cart/empty",
		"/profile/my-order",
		"/profile/my-order/add-review",
		"/profile/my-order/{so_number}",
		"/track/{provider}/{so_number}",
		"/profile/my-order-confirm/{so_number}",
		"/profile/credit-card",
		"/address",
		"/address/add",
		"/address/edit/{id}",
		"/bantuan",
		"/bantuan/{detail}",
	}
)

func (m *App) handler() {
	source := viper.GetString("source")
	m.router.StrictSlash(true)
	// pdp
	m.router.HandleFunc("/{name:[a-z0-9]+}-{id:[0-9]+}.html", reader.PDPHandler).Methods("GET")
	// pcp
	m.router.HandleFunc("/p-{id:[0-9]+}/{name:[a-z0-9]+}", reader.PCPHandler).Methods("GET")
	// brand
	m.router.HandleFunc("/brand/{id:[0-9]+}/{name:[a-z0-9]+}", reader.BrandHandler).Methods("GET")
	// store
	m.router.HandleFunc("/store/{id:[0-9]+}/{name:[a-z0-9]+}", reader.StoreHandler).Methods("GET")
	// campaign
	m.router.HandleFunc("/mau-gaya-itu-gampang/{campaign_name}-{campaign_id}/{post_id}/{icode}", reader.CampaignHandler).Methods("GET")
	// home
	m.router.HandleFunc("/", reader.HomeHandler).Methods("GET")
	// custom url
	for _, name := range customRoutes {
		m.router.HandleFunc(name, reader.CustomHandler).Methods("GET")
	}
	m.router.PathPrefix("/assets/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(source))))
	m.router.PathPrefix("/{name}.js").Handler(http.StripPrefix("/", http.FileServer(http.Dir(source))))
	m.router.PathPrefix("/{name}.css").Handler(http.StripPrefix("/", http.FileServer(http.Dir(source))))
	m.router.PathPrefix("/robots.txt").Handler(http.StripPrefix("/", http.FileServer(http.Dir(source))))
	m.router.NotFoundHandler = http.HandlerFunc(reader.NotFoundHandler)
}

// RevampInit RevampInit init
func (m *App) RevampInit() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	m.router = mux.NewRouter()
	m.handler()
	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", viper.GetString("app.listen")),
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
