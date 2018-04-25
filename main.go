package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/unicolony/hayes/command"
	"github.com/unicolony/hayes/source"
)

func main() {
	source.Init()
	cmd := &cobra.Command{
		Use: "./app",
	}
	cmd.AddCommand(command.ServeCommand())
	cmd.AddCommand(command.RunCommand())
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
	}
}
