package command

import (
	"github.com/spf13/cobra"
)

// RunCommand frontend serve
func RunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "run an app",
		Long:  "run an app",
		Run: func(c *cobra.Command, args []string) {
			c.Help()
		},
	}
	return cmd
}
