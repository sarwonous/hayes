package app

import (
	"./mm"
	"github.com/spf13/cobra"
)

// App App
type App interface {
	Name() string
	Command() *cobra.Command
}

var apps = make(map[string]App, 0)

// Init Init
func Init() {
	addApp(mm.NewMM())
}

func addApp(a App) {
	apps[a.Name()] = a
}

// FindApp FindApp
func FindApp(name string) App {
	return apps[name]
}

// InsertCommand InsertCommand
func InsertCommand(c *cobra.Command) *cobra.Command {
	for _, a := range apps {
		c.AddCommand(a.Command())
	}

	return c
}
