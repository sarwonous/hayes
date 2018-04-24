package command

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/unicolony/hayes/app"
)

// ServeCommand frontend serve
func ServeCommand() *cobra.Command {
	app.Init()
	cmd := &cobra.Command{
		Use:   "serve",
		Short: "serve an app",
		Long:  "serve an app",
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("serve")
		},
	}
	app.InsertCommand(cmd)
	return cmd
}
