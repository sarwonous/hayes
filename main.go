package main

import (
	"fmt"

	"./command"
	"./source"
	"github.com/spf13/cobra"
)

func main() {
	source.Init()
	cmd := &cobra.Command{
		Use: "blog",
	}
	cmd.AddCommand(command.ServeCommand())
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
	}
}
