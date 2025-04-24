package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(CreateCmd)
	RootCmd.AddCommand(ListCmd)
}

var RootCmd = &cobra.Command{
	Use:   "gota",
	Short: "Gota is like make, but written in Go and probably less mature.",
	Long: `Gota is a fast and flexible task runner, project initializer, and lightweight build tool.

It reads tasks from a YAML file in the current working directory and executes them —
kind of like 'make', 'just', or 'task', but with fewer features and more Go.

Includes built-in commands like 'create' to scaffold new projects, so you can spend less time typing boilerplate
and more time pretending to be productive.

Great for automating your dev workflows with minimal guilt.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
