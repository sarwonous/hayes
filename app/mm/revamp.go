package mm

import (
	"context"
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/viper"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/unicolony/hayes/app/mm/reader"
	"github.com/unicolony/hayes/source/log"
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
	reader.SetupRouter(m.router, customRoutes)
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
			log.New().Error("app inited")
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.New().Info("app shutdown")
	os.Exit(0)
}

// Init Init
func (m *App) Init() {
	log.New().Info("app inited")
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
